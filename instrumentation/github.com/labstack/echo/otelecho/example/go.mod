module go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho/example

go 1.15

replace (
	go.opentelemetry.io/contrib => ../../../../../../
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho => ../
	go.opentelemetry.io/contrib/propagators/b3 => ../../../../../../propagators/b3
)

require (
	github.com/labstack/echo/v4 v4.5.0
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.23.0
	go.opentelemetry.io/otel v1.0.1
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC3
	go.opentelemetry.io/otel/sdk v1.0.0-RC3
	go.opentelemetry.io/otel/trace v1.0.1
)
