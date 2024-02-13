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

// Option applies configuration to the [metric.Producer] returned from
// [NewProducer].
type Option interface {
	apply(config) config
}

type config struct {
	filters []func(string) bool
}

func newConfig(opts []Option) config {
	var c config
	for _, o := range opts {
		c = o.apply(c)
	}
	return c
}

type filterOpt struct {
	filter func(string) bool
}

func (o filterOpt) apply(c config) config {
	c.filters = append(c.filters, o.filter)
	return c
}

// WithMetrics returns an Option that will only allow metrics to be generated
// for the provided [runtime.Metric] names. If a metric name is not present it
// will not be used to generate metrics by the configured [metric.Producer]
// returned from [NewProducer].
//
// If multiple of these [Option]s or the [Option]s created by [WithoutMetrics]
// are provided they will be applied in the order they are passed.
func WithMetrics(names ...string) Option {
	if len(names) == 0 {
		return filterOpt{filter: func(string) bool { return false }}
	}

	nameSet := make(map[string]struct{}, len(names))
	for _, n := range names {
		nameSet[n] = struct{}{}
	}
	return filterOpt{
		filter: func(name string) bool {
			_, ok := nameSet[name]
			return ok
		},
	}
}

// WithoutMetrics returns an Option that will not allow metrics to be generated
// for the provided [runtime.Metric] names. If a metric name is not present it
// will continue to be used to generate metrics by the configured
// [metric.Producer] returned from [NewProducer].
//
// If multiple of these [Option]s or the [Option]s created by [WithMetrics] are
// provided they will be applied in the order they are passed.
func WithoutMetrics(names ...string) Option {
	if len(names) == 0 {
		return filterOpt{filter: func(string) bool { return true }}
	}

	nameSet := make(map[string]struct{}, len(names))
	for _, n := range names {
		nameSet[n] = struct{}{}
	}
	return filterOpt{
		filter: func(name string) bool {
			_, ok := nameSet[name]
			return !ok
		},
	}
}
