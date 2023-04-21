# Introduction


# Project Structure
```lua
catalog/
├── cmd/
│   ├── catalog-api/
│   │   └── main.go
│   └── catalog-worker/
│       └── main.go
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   ├── middleware/
│   │   ├── models/
│   │   └── routers/
│   ├── domain/
│   │   ├── entities/
│   │   ├── repositories/
│   │   ├── services/
│   │   └── value_objects/
│   ├── infrastructure/
│   │   ├── database/
│   │   ├── messaging/
│   │   └── storage/
│   └── utils/
├── migrations/
│   └── sql/
├── scripts/
├── test/
├── vendor/
├── Makefile
├── README.md
├── docker-compose.yml
└── go.mod
```
Giải thích chi tiết về từng phần:

* cmd/: Thư mục chứa các file chính của ứng dụng, ví dụ như catalog-api để khởi động HTTP server hoặc catalog-worker để khởi động worker.

* internal/: Thư mục chứa các file và thư mục bên trong của ứng dụng, không được xuất bản bên ngoài gói. Bao gồm các package chứa mã nguồn cho các thành phần của ứng dụng:

** api/: Chứa mã nguồn cho HTTP server và các đối tượng liên quan (handler, middleware, router, model).

** domain/: Chứa các thành phần của lớp domain trong kiến trúc DDD, bao gồm các entities, repositories, services, và value objects.

** infrastructure/: Chứa các thành phần cơ sở hạ tầng của ứng dụng, bao gồm các giao tiếp với cơ sở dữ liệu, bộ nhớ đệm, messaging queue, ...

    *** database: Chứa các tệp cấu hình CSDL, các script tạo CSDL, các tệp liên quan tới ORM.

    *** messaging: Sử dụng để giao tiếp giữa các thành phần khác nhau bên trong microservice hoặc giữa các microservice khác nhau bao gồm: message broker, message queue, message listener, message handler, v.v.

    *** storage: Chứa các struct, interface và phương thức để thực hiện kết nối tới database.

** utils/: Chứa các hàm tiện ích cho ứng dụng.

*    migrations/: Thư mục chứa các tập lệnh để thực hiện cập nhật cơ sở dữ liệu.

*    scripts/: Chứa các tập lệnh hữu ích cho việc triển khai và phát triển ứng dụng.

*    test/: Thư mục chứa các file liên quan đến việc kiểm thử ứng dụng.

*    vendor/: Thư mục chứa các dependency của ứng dụng.

    Makefile: File Makefile chứa các target để dễ dàng build và triển khai ứng dụng.

    README.md: Tài liệu hướng dẫn sử dụng cho ứng dụng.

    docker-compose.yml: File cấu hình docker-compose để khởi động ứng
# Getting Started
This is a overview for using on project
### Build project
1. Docker compose up
```
go build -o catalog-api.exe .\cmd\catalog-api\main.go
```


### Test project
1. To run test at the specific component
```
go test -v .\test\migrations\migrations_test.go
```

2. To run test with the specific case test
```
go test -v .\test\migrations\migrations_test.go -run TestInitCollections
```


# Docker usage

1. We also to use mongodb from community. Refer: [here](https://www.mongodb.com/docs/manual/tutorial/install-mongodb-community-with-docker/)