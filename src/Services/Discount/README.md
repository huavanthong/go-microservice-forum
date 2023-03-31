
# Introduction

# Project Structure
```lua
discount-service/
├── cmd/
│   ├── discount/
│   │   ├── main.go
│   │   └── migrate.go
│   └── coupon/
│       ├── main.go
│       └── migrate.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── controllers/
│   │   ├── discount.go
│   │   └── coupon.go
│   ├── models/
│   │   ├── discount.go
│   │   └── coupon.go
│   ├── repositories/
│   │   ├── discount.go
│   │   └── coupon.go
│   ├── routes/
│   │   ├── discount.go
│   │   └── coupon.go
│   ├── services/
│   │   ├── discount.go
│   │   └── coupon.go
│   ├── utils/
│   │   ├── logger.go
│   │   └── validator.go
│   └── proto/
│       ├── discount/
│       │   ├── discount.pb.go
│       │   └── discount.proto
│       └── coupon/
│           ├── coupon.pb.go
│           └── coupon.proto
├── tests/
│   ├── discount_test.go
│   └── coupon_test.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
└── .env.example
```
Let's go over the different components of this project structure:

* cmd: This folder contains the main applications for the Discount and Coupon microservices. It also contains a migrate.go file that handles database migrations.

* internal: This folder contains the internal packages that the Discount and Coupon microservices use.

* config: This package contains configuration files, such as the database configuration, logging configuration, and other settings.

* controllers: This package contains the controllers for the Discount and Coupon microservices. The controllers are responsible for handling requests and responses.

* models: This package contains the database models for the Discount and Coupon microservices. These models define the schema for the data that the microservices store.

* repositories: This package contains the database repositories for the Discount and Coupon microservices. These repositories are responsible for communicating with the database.

* routes: This package contains the routes for the Discount and Coupon microservices. These routes define the endpoints that the microservices expose.

* services: This package contains the business logic for the Discount and Coupon microservices. The services are responsible for implementing the rules and calculations for determining discounts and coupons.

* utils: This package contains utility functions that are used throughout the Discount and Coupon microservices. For example, it may contain functions for logging or validating data.

* proto: This folder contains the gRPC protocol buffer files for the Discount and Coupon microservices. The .proto files define the messages and services that the microservices use.

* tests: This folder contains the tests for the Discount and Coupon microservices. The tests cover unit testing, integration testing, and end-to-end testing.

* docker-compose.yml: This file contains the configuration for running the Discount and Coupon microservices together using Docker Compose.
