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

package main

import "go.opentelemetry.io/otel/attribute"

type Metric struct {
	Name        string
	Description string
	Unit        string
	Kind        Kind
	Attrs       []Attribute
}

type Attribute struct {
	RuntimeName string
	Attr        []attribute.KeyValue
}

type Kind struct {
	Agg  uint
	Mono bool
	Num  string
}

var (
	KindUndefined              = Kind{}
	KindInt64MonotonicSum      = Kind{Agg: 1, Mono: true, Num: "int64"}
	KindInt64NonMonotonicSum   = Kind{Agg: 1, Mono: false, Num: "int64"}
	KindInt64Gauge             = Kind{Agg: 2, Num: "int64"}
	KindFloat64MonotonicSum    = Kind{Agg: 1, Mono: true, Num: "float64"}
	KindFloat64NonMonotonicSum = Kind{Agg: 1, Mono: false, Num: "float64"}
	KindFloat64Gauge           = Kind{Agg: 2, Num: "float64"}
	KindFloat64Histogram       = Kind{Agg: 3, Num: "float64"}
)

var Metrics = []*Metric{
	{
		Name:        "go.cgo.go_to_c_calls",
		Description: "Count of calls made from Go to C by the current process.",
		Unit:        "{call}",
		Kind:        KindInt64MonotonicSum,
		Attrs: []Attribute{{
			RuntimeName: "/cgo/go-to-c-calls:calls",
			Attr:        []attribute.KeyValue{},
		}},
	},
	{
		Name:        "go.cpu.usage",
		Description: "Estimated CPU time usage. This metric is an overestimate, and not directly comparable to system CPU time measurements. Compare only with other go.cpu.usage metrics.",
		Unit:        "s{cpu}",
		Kind:        KindFloat64MonotonicSum,
		Attrs: []Attribute{{
			RuntimeName: "/cpu/classes/gc/mark/assist:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "gc.mark.assist"),
			},
		}, {
			RuntimeName: "/cpu/classes/gc/mark/dedicated:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "gc.mark.dedicated"),
			},
		}, {
			RuntimeName: "/cpu/classes/gc/mark/idle:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "gc.mark.idle"),
			},
		}, {
			RuntimeName: "/cpu/classes/gc/pause:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "gc.pause"),
			},
		}, {
			RuntimeName: "/cpu/classes/gc/total:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "gc.total"),
			},
		}, {
			RuntimeName: "/cpu/classes/idle:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "idle"),
			},
		}, {
			RuntimeName: "/cpu/classes/scavenge/assist:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "scavenge.assist"),
			},
		}, {
			RuntimeName: "/cpu/classes/scavenge/background:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "scavenge.background"),
			},
		}, {
			RuntimeName: "/cpu/classes/scavenge/total:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "scavenge.total"),
			},
		}, {
			RuntimeName: "/cpu/classes/total:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "total"),
			},
		}, {
			RuntimeName: "/cpu/classes/user:cpu-seconds",
			Attr: []attribute.KeyValue{
				attribute.String("class", "user"),
			},
		}},
	},
	{
		Name:        "go.gc.cycles",
		Description: "Count of completed GC cycles.",
		Unit:        "{cycle}",
		Kind:        KindInt64MonotonicSum,
		Attrs: []Attribute{{
			RuntimeName: "/gc/cycles/automatic:gc-cycles",
			Attr: []attribute.KeyValue{
				attribute.String("trigger", "automatic"),
			},
		}, {
			RuntimeName: "/gc/cycles/forced:gc-cycles",
			Attr: []attribute.KeyValue{
				attribute.String("trigger", "forced"),
			},
		}, {
			RuntimeName: "/gc/cycles/total:gc-cycles",
			Attr:        []attribute.KeyValue{},
		}},
	},
	{
		Name:        "go.gc.gogc",
		Description: "Heap size target percentage configured by the user, otherwise 100. This value is set by the GOGC environment variable, and the runtime/debug.SetGCPercent function.",
		Unit:        "%",
		Kind:        KindInt64Gauge,
		Attrs: []Attribute{{
			RuntimeName: "/gc/gogc:percent",
			Attr:        []attribute.KeyValue{},
		}},
	},
	{
		Name:        "go.gc.gomemlimit",
		Description: "Go runtime memory limit configured by the user, otherwise math.MaxInt64. This value is set by the GOMEMLIMIT environment variable, and the runtime/debug.SetMemoryLimit function.",
		Unit:        "By",
		Kind:        KindInt64Gauge,
		Attrs: []Attribute{{
			RuntimeName: "/gc/gomemlimit:bytes",
			Attr:        []attribute.KeyValue{},
		}},
	},
	{
		Name:        "go.gc.heap.allocs",
		Description: "Distribution of heap allocations by approximate size. Bucket counts increase monotonically. Note that this does not include tiny objects as defined by /gc/heap/tiny/allocs:objects, only tiny blocks.",
		Unit:        "By",
		Kind:        KindFloat64Histogram,
		Attrs: []Attribute{{
			RuntimeName: "/gc/heap/allocs-by-size:bytes",
			Attr:        []attribute.KeyValue{},
		}},
	},
}

/*
var Conversions = map[string]runtime.Conversion{
	"/cgo/go-to-c-calls:calls":                                  Unit: "{calls}"},

	"/cpu/classes/gc/mark/assist:cpu-seconds":                   Unit: "{cpu_seconds}"},
	"/cpu/classes/gc/mark/dedicated:cpu-seconds":                Unit: "{cpu_seconds}"},
	"/cpu/classes/gc/mark/idle:cpu-seconds":                     Unit: "{cpu_seconds}"},
	"/cpu/classes/gc/pause:cpu-seconds":                         Unit: "{cpu_seconds}"},
	"/cpu/classes/gc/total:cpu-seconds":                         Unit: "{cpu_seconds}"},
	"/cpu/classes/idle:cpu-seconds":                             Unit: "{cpu_seconds}"},
	"/cpu/classes/scavenge/assist:cpu-seconds":                  Unit: "{cpu_seconds}"},
	"/cpu/classes/scavenge/background:cpu-seconds":              Unit: "{cpu_seconds}"},
	"/cpu/classes/scavenge/total:cpu-seconds":                   Unit: "{cpu_seconds}"},
	"/cpu/classes/total:cpu-seconds":                            Unit: "{cpu_seconds}"},
	"/cpu/classes/user:cpu-seconds":                             Unit: "{cpu_seconds}"},

	"/gc/cycles/automatic:gc-cycles":                            Unit: "{gc_cycles}"},
	"/gc/cycles/forced:gc-cycles":                               Unit: "{gc_cycles}"},
	"/gc/cycles/total:gc-cycles":                                Unit: "{gc_cycles}"},

	"/gc/gogc:percent":                                          Unit: "{percent}"},

	"/gc/gomemlimit:bytes":                                      Unit: "By"},

	"/gc/heap/allocs:bytes":                                     Unit: "By"},
	"/gc/heap/frees:bytes":                                      Unit: "By"},
	"/gc/heap/goal:bytes":                                       Unit: "By"},
	"/gc/heap/live:bytes":                                       Unit: "By"},

	"/gc/heap/allocs-by-size:bytes":                             Unit: "By"},

	"/gc/heap/frees-by-size:bytes":                              Unit: "By"},

	"/gc/heap/allocs:objects":                                   Unit: "{objects}"},
	"/gc/heap/frees:objects":                                    Unit: "{objects}"},
	"/gc/heap/objects:objects":                                  Unit: "{objects}"},
	"/gc/heap/tiny/allocs:objects":                              Unit: "{objects}"},

	"/gc/limiter/last-enabled:gc-cycle":                         Unit: "{gc_cycle}"},
	"/gc/pauses:seconds":                                        Unit: "s"},
	"/gc/scan/globals:bytes":                                    Unit: "By"},
	"/gc/scan/heap:bytes":                                       Unit: "By"},
	"/gc/scan/stack:bytes":                                      Unit: "By"},
	"/gc/scan/total:bytes":                                      Unit: "By"},
	"/gc/stack/starting-size:bytes":                             Unit: "By"},

	"/godebug/non-default-behavior/execerrdot:events":           Unit: "{events}"},
	"/godebug/non-default-behavior/gocachehash:events":          Unit: "{events}"},
	"/godebug/non-default-behavior/gocachetest:events":          Unit: "{events}"},
	"/godebug/non-default-behavior/gocacheverify:events":        Unit: "{events}"},
	"/godebug/non-default-behavior/gotypesalias:events":         Unit: "{events}"},
	"/godebug/non-default-behavior/http2client:events":          Unit: "{events}"},
	"/godebug/non-default-behavior/http2server:events":          Unit: "{events}"},
	"/godebug/non-default-behavior/httplaxcontentlength:events": Unit: "{events}"},
	"/godebug/non-default-behavior/httpmuxgo121:events":         Unit: "{events}"},
	"/godebug/non-default-behavior/installgoroot:events":        Unit: "{events}"},
	"/godebug/non-default-behavior/jstmpllitinterp:events":      Unit: "{events}"},
	"/godebug/non-default-behavior/multipartmaxheaders:events":  Unit: "{events}"},
	"/godebug/non-default-behavior/multipartmaxparts:events":    Unit: "{events}"},
	"/godebug/non-default-behavior/multipathtcp:events":         Unit: "{events}"},
	"/godebug/non-default-behavior/panicnil:events":             Unit: "{events}"},
	"/godebug/non-default-behavior/randautoseed:events":         Unit: "{events}"},
	"/godebug/non-default-behavior/tarinsecurepath:events":      Unit: "{events}"},
	"/godebug/non-default-behavior/tls10server:events":          Unit: "{events}"},
	"/godebug/non-default-behavior/tlsmaxrsasize:events":        Unit: "{events}"},
	"/godebug/non-default-behavior/tlsrsakex:events":            Unit: "{events}"},
	"/godebug/non-default-behavior/tlsunsafeekm:events":         Unit: "{events}"},
	"/godebug/non-default-behavior/x509sha1:events":             Unit: "{events}"},
	"/godebug/non-default-behavior/x509usefallbackroots:events": Unit: "{events}"},
	"/godebug/non-default-behavior/x509usepolicies:events":      Unit: "{events}"},
	"/godebug/non-default-behavior/zipinsecurepath:events":      Unit: "{events}"},
	"/memory/classes/heap/free:bytes":                           Unit: "By"},
	"/memory/classes/heap/objects:bytes":                        Unit: "By"},
	"/memory/classes/heap/released:bytes":                       Unit: "By"},
	"/memory/classes/heap/stacks:bytes":                         Unit: "By"},
	"/memory/classes/heap/unused:bytes":                         Unit: "By"},
	"/memory/classes/metadata/mcache/free:bytes":                Unit: "By"},
	"/memory/classes/metadata/mcache/inuse:bytes":               Unit: "By"},
	"/memory/classes/metadata/mspan/free:bytes":                 Unit: "By"},
	"/memory/classes/metadata/mspan/inuse:bytes":                Unit: "By"},
	"/memory/classes/metadata/other:bytes":                      Unit: "By"},
	"/memory/classes/os-stacks:bytes":                           Unit: "By"},
	"/memory/classes/other:bytes":                               Unit: "By"},
	"/memory/classes/profiling/buckets:bytes":                   Unit: "By"},
	"/memory/classes/total:bytes":                               Unit: "By"},
	"/sched/gomaxprocs:threads":                                 Unit: "{threads}"},
	"/sched/goroutines:goroutines":                              Unit: "{goroutines}"},
	"/sched/latencies:seconds":                                  Unit: "s"},
	"/sched/pauses/stopping/gc:seconds":                         Unit: "s"},
	"/sched/pauses/stopping/other:seconds":                      Unit: "s"},
	"/sched/pauses/total/gc:seconds":                            Unit: "s"},
	"/sched/pauses/total/other:seconds":                         Unit: "s"},
	"/sync/mutex/wait/total:seconds":                            Unit: "s"},
}
*/
