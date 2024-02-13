// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runtime // import "go.opentelemetry.io/contrib/bridges/runtime"

import (
	"runtime/metrics"
	"time"

	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type Metric struct {
	RuntimeNames []string
	Conversion   Conversion
}

type Conversion struct {
	Name            string
	Description     string
	Unit            string
	AggregationFunc AggregationFunc
}

type AggregationFunc func(*metricdata.Aggregation, time.Time, []*metrics.Sample) error

func reset[T any](slice []T, length, capacity int) []T {
	if cap(slice) < capacity {
		return make([]T, length, capacity)
	}
	return slice[:length]
}

// buckets2Bounds conversts metrics.Float64Histogram buckets and counts to OTel
// bounds and bucket counts.
//
// metrics.Float64Histogram buckets are defined as the inclusive lower-bounds:
//
//	[buckets[i], buckets[i+1]) for 0 < i < len(buckets)-1
//	[buckets[i], +infinity) for i == len(buckets)-1
//	note: len(buckets) != 1 always
//
// OTel bounds are defined as the inclusive upper-bounds:
//
//	(-infinity, bounds[i]] for i == 0
//	(bounds[i-1], bounds[i]] for 0 < i < len(bounds)-1
//	(bounds[i], +infinity) for i == len(bounds)-1
//
// There is a bug in this conversion due to the loss of information in the
// binningprocess. Given the inclusivity of bounds changes from lower-to-higher
// here there are going to be counts on the edges of the bucket ranges that
// should be moved to neighbor bounds. However, we do not know how many of the
// counts to make this conversion to given the lack of origial data. This
// conversion is going to be a best effort.
func buckets2Bounds(buckets []float64, counts []uint64) (bounds []float64, bucketCounts []uint64) {
	switch len(buckets) {
	case 0:
		return bounds, bucketCounts
	case 1:
		// This should never happen, two boundaries are required to describe a
		// bucket. Handle it gracefully anyways if it does.
		return bounds, bucketCounts
	case 2:
		bounds = []float64{buckets[1]}
		bucketCounts = []uint64{counts[0], 0}
	default:
		bounds = make([]float64, len(buckets)-2)
		copy(bounds, buckets[1:len(buckets)-1])
		bucketCounts = counts
	}

	return bounds, bucketCounts
}

func sum(vals []uint64) (ans uint64) {
	for _, v := range vals {
		ans += v
	}
	return ans
}
