module go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful

go 1.15

replace (
	go.opentelemetry.io/contrib => ../../../../../
	go.opentelemetry.io/contrib/propagators/b3 => ../../../../../propagators/b3
)

require (
	github.com/emicklei/go-restful/v3 v3.5.2
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/contrib v0.23.0
	go.opentelemetry.io/contrib/propagators/b3 v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC3
	go.opentelemetry.io/otel/trace v1.0.0-RC3
)
