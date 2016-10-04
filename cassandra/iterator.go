// Copyright 2016 The Vulcan Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cassandra

import (
	"github.com/digitalocean/vulcan/convert"
	"github.com/gocql/gocql"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/storage/local"
	"github.com/prometheus/prometheus/storage/metric"
)

const fetchUncompressedSQLIter = `SELECT at, value FROM uncompressed WHERE fqmn = ? AND at >= ? AND at <= ? ORDER BY at ASC`

// SeriesIterator is a Cassandra-backed implementation of a prometheus SeriesIterator.
type SeriesIterator struct {
	iter       *gocql.Iter
	m          metric.Metric
	curr, last *model.SamplePair
	ready      chan struct{}
}

// SeriesIteratorConfig is used in NewSeriesIterator to create a SeriesIterator.
type SeriesIteratorConfig struct {
	Session       *gocql.Session
	Metric        metric.Metric
	After, Before model.Time
	PageSize      int
	Prefetch      float64
}

// NewSeriesIterator creates a Cassandra-backed implementation of a prometheus storage
// SeriesIterator. This iterator immediately begins pre-fetching data upon creation.
func NewSeriesIterator(config *SeriesIteratorConfig) *SeriesIterator {
	ts := convert.MetricToTimeSeries(config.Metric)
	fqmn := ts.ID()
	si := &SeriesIterator{
		m: config.Metric,
		curr: &model.SamplePair{
			Timestamp: local.ZeroSamplePair.Timestamp,
			Value:     local.ZeroSamplePair.Value,
		},
		last: &model.SamplePair{
			Timestamp: local.ZeroSamplePair.Timestamp,
			Value:     local.ZeroSamplePair.Value,
		},
		ready: make(chan struct{}),
	}
	// creating a gocql iterator takes time, so we instantiate it inside of a goroutine so
	// we can return quickly from NewSeriesIterator which is important for the performance
	// of the prometheus query engine. The si.ready channel signals when the iter is ready
	// to be used.
	go func() {
		si.iter = config.Session.Query(fetchUncompressedSQLIter, fqmn, config.After, config.Before).
			PageSize(config.PageSize).
			Prefetch(config.Prefetch).
			Iter()
		close(si.ready)
	}()
	return si
}

// ValueAtOrBeforeTime gets the value that is closest before the given time. In case a value
// exists at precisely the given time, that value is returned. If no
// applicable value exists, ZeroSamplePair is returned. This function assumes that
// ValueAtOrBeforeTime will be called only with incrementing values of t and that
// this SeriesIterator will only call either ValueAtOrBeforeTime or RangeValues, but
// not both functions.
func (si *SeriesIterator) ValueAtOrBeforeTime(t model.Time) model.SamplePair {
	<-si.ready
	// curr == nil means that there are no more values to iterate over.
	if si.curr == nil {
		return *si.last
	}
	for {
		if si.curr.Timestamp > t {
			return *si.last
		}
		si.last.Timestamp = si.curr.Timestamp
		si.last.Value = si.curr.Value
		ok := si.iter.Scan(&si.curr.Timestamp, &si.curr.Value)
		if !ok {
			// done iterating; set curr to nil to signal no more values on iter.
			si.curr = nil
			return *si.last
		}
	}
}

// RangeValues gets all values contained within a given interval. RangeValues assumes
// that the interval values OldestInclusive and NewestInclusive will always be
// higher than the previous call to RangeValues.
func (si *SeriesIterator) RangeValues(r metric.Interval) []model.SamplePair {
	<-si.ready
	result := []model.SamplePair{}
	// curr == nil means that there are no more values to iterate over.
	if si.curr == nil {
		return result
	}
	for {
		// the iterator has advanced past our current interval.
		if si.curr.Timestamp > r.NewestInclusive {
			return result
		}
		// add curr from previous run to result, but exclude starting local.ZeroValuePair.
		if si.curr.Timestamp >= r.OldestInclusive {
			result = append(result, model.SamplePair{
				Timestamp: si.curr.Timestamp,
				Value:     si.curr.Value,
			})
		}
		// if we are exactly at the upper time bound, return result and DO NOT advance the iterator
		// so that this value can also be added to the next call to RangeValues as its lower time bound
		// value. This assumes that there will be no two samples at the same timestamp (enforced by cassandra schema).
		if si.curr.Timestamp == r.NewestInclusive {
			return result
		}
		ok := si.iter.Scan(&si.curr.Timestamp, &si.curr.Value)
		if !ok {
			si.curr = nil
			return result
		}
	}
}

// Metric returns the metric of the series that the iterator corresponds to.
func (si *SeriesIterator) Metric() metric.Metric {
	return si.m
}

// Close closes the iterator and releases the underlying data.
func (si *SeriesIterator) Close() {
	<-si.ready
	err := si.iter.Close()
	if err != nil {
		panic(err)
	}
	return
}
