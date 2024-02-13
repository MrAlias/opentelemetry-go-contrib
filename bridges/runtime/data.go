// Code created by datagen. DO NOT MODIFY.
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
	"errors"
	"fmt"
	"runtime/metrics"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

var All = []Metric{
	{
		RuntimeNames: []string{
			"/cgo/go-to-c-calls:calls",
		},
		Conversion: Conversion{
			Name:        "go.cgo.go_to_c_calls",
			Description: "Count of calls made from Go to C by the current process.",
			Unit:        "{call}",
			AggregationFunc: func() AggregationFunc {
				attrMap := map[string]attribute.Set{
					"/cgo/go-to-c-calls:calls": attribute.NewSet(),
				}
				return func(agg *metricdata.Aggregation, t time.Time, samples []*metrics.Sample) error {
					if agg == nil {
						return errors.New("nil aggregation")
					}
					var err error
					sum, _ := (*agg).(metricdata.Sum[int64])
					sum.Temporality = metricdata.CumulativeTemporality
					sum.IsMonotonic = true
					sum.DataPoints = reset(sum.DataPoints, 0, len(samples))
					for _, s := range samples {
						a, ok := attrMap[s.Name]
						if !ok {
							err = errors.Join(err, fmt.Errorf(
								"missing attribute: %s", s.Name,
							))
						}
						sum.DataPoints = append(sum.DataPoints, metricdata.DataPoint[int64]{
							StartTime:  start,
							Time:       t,
							Attributes: a,
							Value:      int64(s.Value.Uint64()),
						})
					}
					*agg = sum
					return err
				}
			}(),
		},
	},
	{
		RuntimeNames: []string{
			"/cpu/classes/gc/mark/assist:cpu-seconds",
			"/cpu/classes/gc/mark/dedicated:cpu-seconds",
			"/cpu/classes/gc/mark/idle:cpu-seconds",
			"/cpu/classes/gc/pause:cpu-seconds",
			"/cpu/classes/gc/total:cpu-seconds",
			"/cpu/classes/idle:cpu-seconds",
			"/cpu/classes/scavenge/assist:cpu-seconds",
			"/cpu/classes/scavenge/background:cpu-seconds",
			"/cpu/classes/scavenge/total:cpu-seconds",
			"/cpu/classes/total:cpu-seconds",
			"/cpu/classes/user:cpu-seconds",
		},
		Conversion: Conversion{
			Name:        "go.cpu.usage",
			Description: "Estimated CPU time usage. This metric is an overestimate, and not directly comparable to system CPU time measurements. Compare only with other go.cpu.usage metrics.",
			Unit:        "s{cpu}",
			AggregationFunc: func() AggregationFunc {
				attrMap := map[string]attribute.Set{
					"/cpu/classes/gc/mark/assist:cpu-seconds": attribute.NewSet(
						attribute.String("class", "gc.mark.assist"),
					),
					"/cpu/classes/gc/mark/dedicated:cpu-seconds": attribute.NewSet(
						attribute.String("class", "gc.mark.dedicated"),
					),
					"/cpu/classes/gc/mark/idle:cpu-seconds": attribute.NewSet(
						attribute.String("class", "gc.mark.idle"),
					),
					"/cpu/classes/gc/pause:cpu-seconds": attribute.NewSet(
						attribute.String("class", "gc.pause"),
					),
					"/cpu/classes/gc/total:cpu-seconds": attribute.NewSet(
						attribute.String("class", "gc.total"),
					),
					"/cpu/classes/idle:cpu-seconds": attribute.NewSet(
						attribute.String("class", "idle"),
					),
					"/cpu/classes/scavenge/assist:cpu-seconds": attribute.NewSet(
						attribute.String("class", "scavenge.assist"),
					),
					"/cpu/classes/scavenge/background:cpu-seconds": attribute.NewSet(
						attribute.String("class", "scavenge.background"),
					),
					"/cpu/classes/scavenge/total:cpu-seconds": attribute.NewSet(
						attribute.String("class", "scavenge.total"),
					),
					"/cpu/classes/total:cpu-seconds": attribute.NewSet(
						attribute.String("class", "total"),
					),
					"/cpu/classes/user:cpu-seconds": attribute.NewSet(
						attribute.String("class", "user"),
					),
				}
				return func(agg *metricdata.Aggregation, t time.Time, samples []*metrics.Sample) error {
					if agg == nil {
						return errors.New("nil aggregation")
					}
					var err error
					sum, _ := (*agg).(metricdata.Sum[float64])
					sum.Temporality = metricdata.CumulativeTemporality
					sum.IsMonotonic = true
					sum.DataPoints = reset(sum.DataPoints, 0, len(samples))
					for _, s := range samples {
						a, ok := attrMap[s.Name]
						if !ok {
							err = errors.Join(err, fmt.Errorf(
								"missing attribute: %s", s.Name,
							))
						}
						sum.DataPoints = append(sum.DataPoints, metricdata.DataPoint[float64]{
							StartTime:  start,
							Time:       t,
							Attributes: a,
							Value:      s.Value.Float64(),
						})
					}
					*agg = sum
					return err
				}
			}(),
		},
	},
	{
		RuntimeNames: []string{
			"/gc/cycles/automatic:gc-cycles",
			"/gc/cycles/forced:gc-cycles",
			"/gc/cycles/total:gc-cycles",
		},
		Conversion: Conversion{
			Name:        "go.gc.cycles",
			Description: "Count of completed GC cycles.",
			Unit:        "{cycle}",
			AggregationFunc: func() AggregationFunc {
				attrMap := map[string]attribute.Set{
					"/gc/cycles/automatic:gc-cycles": attribute.NewSet(
						attribute.String("trigger", "automatic"),
					),
					"/gc/cycles/forced:gc-cycles": attribute.NewSet(
						attribute.String("trigger", "forced"),
					),
					"/gc/cycles/total:gc-cycles": attribute.NewSet(),
				}
				return func(agg *metricdata.Aggregation, t time.Time, samples []*metrics.Sample) error {
					if agg == nil {
						return errors.New("nil aggregation")
					}
					var err error
					sum, _ := (*agg).(metricdata.Sum[int64])
					sum.Temporality = metricdata.CumulativeTemporality
					sum.IsMonotonic = true
					sum.DataPoints = reset(sum.DataPoints, 0, len(samples))
					for _, s := range samples {
						a, ok := attrMap[s.Name]
						if !ok {
							err = errors.Join(err, fmt.Errorf(
								"missing attribute: %s", s.Name,
							))
						}
						sum.DataPoints = append(sum.DataPoints, metricdata.DataPoint[int64]{
							StartTime:  start,
							Time:       t,
							Attributes: a,
							Value:      int64(s.Value.Uint64()),
						})
					}
					*agg = sum
					return err
				}
			}(),
		},
	},
	{
		RuntimeNames: []string{
			"/gc/gogc:percent",
		},
		Conversion: Conversion{
			Name:        "go.gc.gogc",
			Description: "Heap size target percentage configured by the user, otherwise 100. This value is set by the GOGC environment variable, and the runtime/debug.SetGCPercent function.",
			Unit:        "%",
			AggregationFunc: func() AggregationFunc {
				attrMap := map[string]attribute.Set{
					"/gc/gogc:percent": attribute.NewSet(),
				}
				return func(agg *metricdata.Aggregation, t time.Time, samples []*metrics.Sample) error {
					if agg == nil {
						return errors.New("nil aggregation")
					}
					var err error
					gauge, _ := (*agg).(metricdata.Gauge[int64])
					gauge.DataPoints = reset(gauge.DataPoints, 0, len(samples))
					for _, s := range samples {
						a, ok := attrMap[s.Name]
						if !ok {
							err = errors.Join(err, fmt.Errorf(
								"missing attribute: %s", s.Name,
							))
						}
						gauge.DataPoints = append(gauge.DataPoints, metricdata.DataPoint[int64]{
							Time:       t,
							Attributes: a,
							Value:      int64(s.Value.Uint64()),
						})
					}
					*agg = gauge
					return err
				}
			}(),
		},
	},
	{
		RuntimeNames: []string{
			"/gc/gomemlimit:bytes",
		},
		Conversion: Conversion{
			Name:        "go.gc.gomemlimit",
			Description: "Go runtime memory limit configured by the user, otherwise math.MaxInt64. This value is set by the GOMEMLIMIT environment variable, and the runtime/debug.SetMemoryLimit function.",
			Unit:        "By",
			AggregationFunc: func() AggregationFunc {
				attrMap := map[string]attribute.Set{
					"/gc/gomemlimit:bytes": attribute.NewSet(),
				}
				return func(agg *metricdata.Aggregation, t time.Time, samples []*metrics.Sample) error {
					if agg == nil {
						return errors.New("nil aggregation")
					}
					var err error
					gauge, _ := (*agg).(metricdata.Gauge[int64])
					gauge.DataPoints = reset(gauge.DataPoints, 0, len(samples))
					for _, s := range samples {
						a, ok := attrMap[s.Name]
						if !ok {
							err = errors.Join(err, fmt.Errorf(
								"missing attribute: %s", s.Name,
							))
						}
						gauge.DataPoints = append(gauge.DataPoints, metricdata.DataPoint[int64]{
							Time:       t,
							Attributes: a,
							Value:      int64(s.Value.Uint64()),
						})
					}
					*agg = gauge
					return err
				}
			}(),
		},
	},
	{
		RuntimeNames: []string{
			"/gc/heap/allocs-by-size:bytes",
		},
		Conversion: Conversion{
			Name:        "go.gc.heap.allocs",
			Description: "Distribution of heap allocations by approximate size. Bucket counts increase monotonically. Note that this does not include tiny objects as defined by /gc/heap/tiny/allocs:objects, only tiny blocks.",
			Unit:        "By",
			AggregationFunc: func() AggregationFunc {
				attrMap := map[string]attribute.Set{
					"/gc/heap/allocs-by-size:bytes": attribute.NewSet(),
				}
				return func(agg *metricdata.Aggregation, t time.Time, samples []*metrics.Sample) error {
					if agg == nil {
						return errors.New("nil aggregation")
					}
					var err error
					hist, _ := (*agg).(metricdata.Histogram[float64])
					hist.Temporality = metricdata.CumulativeTemporality
					hist.DataPoints = reset(hist.DataPoints, 0, len(samples))
					for _, s := range samples {
						a, ok := attrMap[s.Name]
						if !ok {
							err = errors.Join(err, fmt.Errorf(
								"missing attribute: %s", s.Name,
							))
						}
						h := s.Value.Float64Histogram()
						bounds, bucketCounts := buckets2Bounds(h.Buckets, h.Counts)
						hist.DataPoints = append(hist.DataPoints, metricdata.HistogramDataPoint[float64]{
							StartTime:    start,
							Time:         t,
							Attributes:   a,
							Count:        sum(h.Counts),
							Bounds:       bounds,
							BucketCounts: bucketCounts,
						})
					}
					*agg = hist
					return err
				}
			}(),
		},
	},
}
