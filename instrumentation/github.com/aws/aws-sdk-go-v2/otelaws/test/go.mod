module go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws/test

go 1.19

require (
	github.com/aws/aws-sdk-go-v2 v1.21.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.21.5
	github.com/aws/aws-sdk-go-v2/service/route53 v1.30.1
	github.com/aws/smithy-go v1.15.0
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.45.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.opentelemetry.io/otel/trace v1.19.0
)

require (
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.42 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.36 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.14 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.35 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v1.24.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws => ../
