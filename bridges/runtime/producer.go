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

//go:generate go run ./internal/datagen/... --out=./data.go

import (
	"context"
	"errors"
	"runtime/metrics"
	"time"

	"go.opentelemetry.io/otel/sdk/instrumentation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

// Scope is the instrumentation scope of the telemetry this package provides.
var Scope = instrumentation.Scope{
	Name:    "go.opentelemetry.io/otel/bridge/runtime",
	Version: "0.0.1",
}

var (
	// now is used for overrides with testing
	now = time.Now
	// start is the start time for accumulation of telemetry.
	start = now()
)

// NewProducer returns a [metric.Producer] that produces metrics from
// [runtime/metrics].
func NewProducer(options ...Option) metric.Producer {
	filters := newConfig(options).filters

	all := All

	var sampleN int
	for _, m := range All {
		if len(filters) > 0 {
			filtered := m.RuntimeNames[:0] // Use the same underlying storage.
			for _, f := range filters {
				for _, name := range m.RuntimeNames {
					if f(name) {
						filtered = append(filtered, name)
					}
				}
			}
			m.RuntimeNames = filtered
		}
		sampleN += len(m.RuntimeNames)
	}

	p := producer{
		sm: []metricdata.ScopeMetrics{{
			Scope:   Scope,
			Metrics: make([]metricdata.Metrics, len(all)),
		}},
		samples:    make([]metrics.Sample, 0, sampleN),
		processors: make([]processor, len(all)),
	}

	for i, m := range all {
		if len(m.RuntimeNames) == 0 {
			// All names filtered out.
			continue
		}

		p.sm[0].Metrics[i] = metricdata.Metrics{
			Name:        m.Conversion.Name,
			Description: m.Conversion.Description,
			Unit:        m.Conversion.Unit,
		}

		samples := make([]*metrics.Sample, len(m.RuntimeNames))
		for i, name := range m.RuntimeNames {
			p.samples = append(p.samples, metrics.Sample{Name: name})
			samples[i] = &p.samples[len(p.samples)-1]
		}

		p.processors[i] = processor{
			Data:    &p.sm[0].Metrics[i].Data,
			Samples: samples,
			Func:    m.Conversion.AggregationFunc,
		}
	}
	return p
}

type processor struct {
	Data    *metricdata.Aggregation
	Samples []*metrics.Sample
	Func    AggregationFunc
}

type producer struct {
	sm         []metricdata.ScopeMetrics
	samples    []metrics.Sample
	processors []processor
}

func (p producer) Produce(context.Context) ([]metricdata.ScopeMetrics, error) {
	t := now()
	metrics.Read(p.samples)

	var err error
	for _, proc := range p.processors {
		errors.Join(err, proc.Func(proc.Data, t, proc.Samples))
	}
	return p.sm, err
}
