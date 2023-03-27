# Introduction

# Project Structure
```lua
.
├── cmd
│   └── basket
│       └── main.go
├── internal
│   ├── domain
│   │   ├── entities
│   │   │   ├── basket.go
│   │   │   └── item.go
│   │   ├── repositories
│   │   │   └── basket_repository.go
│   │   └── services
│   │       ├── basket_service.go
│   │       └── item_service.go
│   ├── infrastructure
│   │   ├── config
│   │   │   ├── config.go
│   │   │   └── config.yaml
│   │   ├── database
│   │   │   ├── database.go
│   │   │   └── migrations
│   │   │       ├── 1_create_basket_table.up.sql
│   │   │       ├── 1_create_basket_table.down.sql
│   │   │       ├── 2_create_item_table.up.sql
│   │   │       └── 2_create_item_table.down.sql
│   │   ├── logging
│   │   │   ├── logger.go
│   │   │   └── logger_test.go
│   │   ├── persistence
│   │   │   ├── mongodb
│   │   │   │   ├── mongodb_client.go
│   │   │   │   └── mongodb_repository.go
│   │   │   └── repository.go
│   │   └── redis
│   │       └── redis_repository.go
│   └── interfaces
│       ├── api
│       │   ├── controllers
│       │   │   ├── basket_controller.go
│       │   │   └── item_controller.go
│       │   ├── middleware
│       │   │   ├── authentication.go
│       │   │   └── logging.go
│       │   ├── routes
│       │   │   └── basket_routes.go
│       │   └── server.go
│       └── persistence
│           └── repository.go
├── Dockerfile
└── go.mod
```