module go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux/example

go 1.20

replace go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux => ../

require (
	github.com/gorilla/mux v1.8.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.45.0
	go.opentelemetry.io/otel v1.20.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.20.0
	go.opentelemetry.io/otel/sdk v1.20.0
	go.opentelemetry.io/otel/trace v1.20.0
)

require (
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
)
