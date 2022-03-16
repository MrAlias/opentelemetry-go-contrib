// Deprecated: This module is no longer supported. Use the
// go.opentelemetry.io/otel/exporters/otlp/otlpmetric exporter as a replacement
// to send data to a collector which can then export with its PRW exporter.
module go.opentelemetry.io/contrib/exporters/metric/cortex

go 1.16

require (
	github.com/golang/snappy v0.0.4
	github.com/google/go-cmp v0.5.7
	// Note: v1.8.2-0.20210928085443-fafb309d4027 is
	// Prometheus v2.30.1 released 2021-09-28
	// https://github.com/prometheus/prometheus/commit/fafb309d4027b050c917362d7d2680c5ad6f6e9e
	github.com/prometheus/prometheus v1.8.2-0.20210928085443-fafb309d4027
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/otel v1.5.0
	go.opentelemetry.io/otel/metric v0.27.0
	go.opentelemetry.io/otel/sdk v1.5.0
	go.opentelemetry.io/otel/sdk/metric v0.27.0
)
