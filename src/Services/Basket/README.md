# Table of Content
* [Introduction](#introduction)
* [Project Structure](#project-structure)
* [Architecture overall](#architecture-overall)
* [Getting Started](#getting-started)

## Introduction
Basket microservice is a component of a larger system that focuses on managing the functionality related to the shopping basket or cart.  
Implemented using the Go programming language, this microservice handles operations such as:
* Adding items to the basket, 
* Removing items, 
* Updating quantities, 
* Retrieving basket information. 
It serves as a crucial part of e-commerce platforms or any application that involves shopping functionality. Leveraging the power and efficiency of Go, the Basket microservice provides a scalable and robust solution for managing customer shopping carts in a fast and reliable manner.

## Project Structure
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

* domain/: Chứa các thành phần của lớp domain trong kiến trúc DDD, bao gồm các entities, repositories, services, và value objects.
* entities: định nghĩa các struct tượng trưng cho các đối tượng trong domain.
* repositories: chứa các interface và implementation của repository pattern, định nghĩa các phương thức để truy vấn và lưu trữ dữ liệu của domain.
* services: chứa các business logic của domain, tương tác với các repositories để thực hiện các thao tác CRUD và xử lý các luồng logic phức tạp hơn.

* infrastructure/: Chứa các phần mềm và công nghệ cơ bản cho project, như cấu hình, logging, database, cache, messaging...
* config: chứa cấu hình của ứng dụng, đọc và parse từ file config.yaml để sử dụng trong runtime.
* database: chứa module cho database, có thể là SQL, NoSQL hay Graph Database, đi kèm với đó là các migration script để khởi tạo schema và seed data cho database.
* logging: chứa module cho logging, cung cấp các hàm log.Error(), log.Warning(), log.Info() để ghi lại các thông tin và lỗi trong quá trình chạy ứng dụng.
* persistence: chứa các implementation cho các interface của repository pattern, xử lý các thao tác lưu trữ dữ liệu với các công nghệ khác nhau, như MongoDB hay Redis.

* interfaces/: chứa các API endpoints và web server, tương tác với services để xử lý các request và response.

* api: chứa các file controller và middleware cho RESTful API, định nghĩa các endpoint và xử lý các input và output.

* controllers:
* persistence: chứa các interface cho repository pattern, làm trung gian giữa services và các implementation của persistence.
* utils/: Chứa các hàm tiện ích cho ứng dụng.
* test/: Thư mục chứa các file liên quan đến việc kiểm thử ứng dụng.
* Makefile: File Makefile chứa các target để dễ dàng build và triển khai ứng dụng.
* README.md: Tài liệu hướng dẫn sử dụng cho ứng dụng.
* docker-compose.yml: File cấu hình docker-compose để khởi động ứng
## Architecture overall
Basket microservice includes 3 service:
* **basketdb**: It is a MongoDB database for storing data on basket microservice.
* **basketredis**: It is a Redis Caching for storing data in caching, help your data can access quickly.
* **basket.api**: It is a backend server to implement RESTful API for basket microservice.
## Getting Started
There are several ways to run the basket microservice for this project:
1. **Docker Compose**: 
    * You can run the microservice by utilizing Docker Compose. It allows you to combine all the services in a docker-compose file and run them together as a unified stack.
    * More details: [Method 1: Docker Compose](#method-1-docker-compose)
2. **Local Machine**: 
    * In certain cases, if the basket API service fails, you might need to run MongoDB locally. This ensures that the necessary database is available for the microservice to function properly.
    * More details: [Method 2: Local Machine](#method-2-local-machine)


### Method 1: Docker Compose
To rebuild image for testing
```
docker-compose -f docker-compose.yml -f docker-compose.override.yml up --build
```
In case if you fail any microservice inside, you can build and run it by docker command.
#### Basket service
Build basket microservice image:
```
docker build -t basket-service .
```

After build successully, we can run container by command:
```
docker run -p 8001:80 --env MONGODB_LOCAL_URI=mongodb://root:password123@basketdb:27018 -it --rm basketapi

docker run -p 8080:8080 --env MONGODB_LOCAL_URI=mongodb://mongodb:27017/basketdb REDIS_URL=localhost:6379 basket-service
```
### Method 2: Local Machine
In case basket.api microservice fail, you need to debug and test on local machine, you can follow the below steps to conduct testing basket.api for local machine.  
**Note:**
    * Make sure basketdb and basketredis container running on docker image.
### Build basket microservice in local machine
1. Buidl app
```
go build -o basket-api.exe .\cmd\basket-api\main.go
```

2. Run app
```
./basket-api.exe
```
### Build MongoDB Database on local machine.
```
docker run -d \
  --name basketdb \
  --restart always \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=password123 \
  -e MONGODB_LOCAL_URI=mongodb://root:password123@basketdb:27017 \
  -p 27017:27017 \
  --expose 27017 \
  -v mongo_data_basket:/data/db \
  mongo

```
#### Redis 
```
docker run -d \
  --name basketredis \
  --restart always \
  -e REDIS_URL=basketredis:6379 \
  -e REDIS_HOST=basketredis \
  -e REDIS_PORT=6379 \
  -e REDIS_PASSWORD=eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81 \
  -p 6379:6379 \
  --expose 6379 \
  -v redis_cache_basket:/data/redis \
  redis \
  redis-server --save 60 1 --loglevel warning --requirepass eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81

```

#### Basket API
```
docker run -d \
  --name basket.api \
  --restart always \
  --link basketdb:basketdb \
  --link basketredis:basketredis \
  --env-file ./Services/Basket/internal/infrastructure/config/app.env \
  -p 8001:8001 \
  basket.api
```
## Testing
To build test at the specific component
```
go test -v .\test\internal\domain\repositories\Basket.repository_test.go
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

