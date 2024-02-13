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
	"context"
	"testing"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

func TestNewMetricProducer(t *testing.T) {
	p := NewProducer()
	r := metric.NewManualReader(metric.WithProducer(p))
	mp := metric.NewMeterProvider(metric.WithReader(r))
	ctx := context.Background()
	t.Cleanup(func() { _ = mp.Shutdown(ctx) })

	rm := new(metricdata.ResourceMetrics)
	r.Collect(ctx, rm)

	for _, sm := range rm.ScopeMetrics {
		t.Logf("Scope: %s, Version: %s", sm.Scope.Name, sm.Scope.Version)
		for _, m := range sm.Metrics {
			t.Logf("- %s (%v): %s", m.Name, m.Unit, m.Description)
			t.Log(m.Data)
		}
	}
	t.Fail()
}
