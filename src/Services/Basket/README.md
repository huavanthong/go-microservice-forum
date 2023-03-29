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
│   │   │   │  
│   │       └── redis
│   │           └── redis_repository.go
│   │      
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
Explain project structure:

* cmd/: Thư mục chứa các file chính của ứng dụng, ví dụ như basket để khởi động HTTP server.

* internal/: Thư mục chứa các file và thư mục bên trong của ứng dụng, không được xuất bản bên ngoài gói. Bao gồm các package chứa mã nguồn cho các thành phần của ứng dụng:

** domain/: Chứa các thành phần của lớp domain trong kiến trúc DDD, bao gồm các entities, repositories, services, và value objects.

** infrastructure/: Chứa các thành phần cơ sở hạ tầng của ứng dụng, bao gồm các giao tiếp với cơ sở dữ liệu, bộ nhớ đệm, messaging queue, ...

    *** database: typically contains modules for managing database connections, such as database configuration files, migration files, and database client libraries. It's responsible for initializing the connection to the database and handling any required migrations.

    *** logging: typically containes logger 

    *** persistence: typically contains modules that define the repository interface and provide the actual implementation of methods that interact with the database. 

** interfaces/: Chứa mã nguồn cho HTTP server và các đối tượng liên quan (handler, middleware, router, model).


** utils/: Chứa các hàm tiện ích cho ứng dụng.

* test/: Thư mục chứa các file liên quan đến việc kiểm thử ứng dụng.

Makefile: File Makefile chứa các target để dễ dàng build và triển khai ứng dụng.

README.md: Tài liệu hướng dẫn sử dụng cho ứng dụng.

docker-compose.yml: File cấu hình docker-compose để khởi động ứng


# Getting Started
## Basket service
Build basket microservice image:
```
docker build -t basket-service .
```

After build successully, we can run container by command:
```
docker run -p 8080:8080 --env MONGODB_LOCAL_URI=mongodb://mongodb:27017/basketdb REDIS_URL=localhost:6379 basket-service
```

## Redis 