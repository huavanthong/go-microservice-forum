module github.com/huavanthong/microservice-golang/product-api

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
	github.com/hashicorp/go-hclog v0.12.1
	github.com/huavanthong/microservice-golang v0.0.0-20220507071640-37bc3cd59bc1 // indirect
	github.com/huavanthong/microservice-golang/currency v0.0.0-20220507034548-1beb3ecf07a1
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/nicholasjackson/building-microservices-youtube/product-api v0.0.0-20211011132451-0cd586295712
	github.com/nicholasjackson/env v0.6.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.46.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace github.com/huavanthong/microservice-golang/currency => ../currency
