module github.com/huavanthong/microservice-golang/product-api-v3

go 1.14

require (
	github.com/go-openapi/errors v0.20.2
	github.com/go-openapi/runtime v0.23.3
	github.com/go-openapi/strfmt v0.21.2
	github.com/go-openapi/swag v0.21.1
	github.com/go-openapi/validate v0.21.0
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.8.0
	github.com/hashicorp/go-hclog v1.2.2
	github.com/huavanthong/microservice-golang/currency v0.0.0-20220507034548-1beb3ecf07a1
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/nicholasjackson/env v0.6.0
	github.com/stretchr/testify v1.7.2
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	google.golang.org/grpc v1.46.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/huavanthong/microservice-golang/currency => ../currency
