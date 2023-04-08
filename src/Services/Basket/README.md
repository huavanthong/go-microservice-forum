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
│       │   ├── models
│       │   │   ├── request
│       │   │   │   ├── basket_request.go
│       │   │   │   └── item_request.go
│       │   │   └── response
│       │   │       ├── basket_response.go
│       │   │       └── item_response.go
│       │   ├── routes
│       │   │   └── basket_routes.go
│       │   └── server.go
│       └── persistence
│           └── repository.go
├── Dockerfile
└── go.mod
```
Explain project structure:

* cmd/: Folder chứa các file main.go để chạy server hoặc worker, đây là entrypoint của chương trình.

* internal/: Folder chính của project, chứa toàn bộ mã nguồn của project, được chia thành các lớp và module theo kiến trúc Clean Architecture.

** domain/: Chứa các thành phần của lớp domain trong kiến trúc DDD, bao gồm các entities, repositories, services, và value objects.
    
    *** entities: định nghĩa các struct tượng trưng cho các đối tượng trong domain.

    *** repositories: chứa các interface và implementation của repository pattern, định nghĩa các phương thức để truy vấn và lưu trữ dữ liệu của domain.

    *** services: chứa các business logic của domain, tương tác với các repositories để thực hiện các thao tác CRUD và xử lý các luồng logic phức tạp hơn.

** infrastructure/: Chứa các phần mềm và công nghệ cơ bản cho project, như cấu hình, logging, database, cache, messaging...

    *** config: chứa cấu hình của ứng dụng, đọc và parse từ file config.yaml để sử dụng trong runtime.

    *** database: chứa module cho database, có thể là SQL, NoSQL hay Graph Database, đi kèm với đó là các migration script để khởi tạo schema và seed data cho database.

    *** logging: chứa module cho logging, cung cấp các hàm log.Error(), log.Warning(), log.Info() để ghi lại các thông tin và lỗi trong quá trình chạy ứng dụng.

    *** persistence: chứa các implementation cho các interface của repository pattern, xử lý các thao tác lưu trữ dữ liệu với các công nghệ khác nhau, như MongoDB hay Redis.

** interfaces/: chứa các API endpoints và web server, tương tác với services để xử lý các request và response.

    *** api: chứa các file controller và middleware cho RESTful API, định nghĩa các endpoint và xử lý các input và output.

        **** controllers:

    *** persistence: chứa các interface cho repository pattern, làm trung gian giữa services và các implementation của persistence.


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

## Running app on local
1. Buidl app
```
go build -o basket-api.exe .\cmd\basket-api\main.go
```

2. Run app
```
./basket-api.exe
```

## Swagger
1. Install swag
```
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate documents for swagger
```
swag init -g ./cmd/basket-api/main.go --output docs
```

3. Access swagger on basket microservice
```
http://localhost:8001/api/v1/swagger/index.html
```