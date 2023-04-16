# Introduction


# Project Structure
```lua 
order-microservice/
├── cmd/
│   ├── order-api/
│   │   ├── main.go
│   │   ├── config.go
│   │   └── ...
│   └── order-worker/
│       ├── main.go
│       ├── config.go
│       └── ...
├── internal/
│   ├── api/
│   │   ├── controllers/
│   │   │   └── order_controller.go
│   │   ├── middleware
│   │   │   ├── authentication.go
│   │   │   ├── cors_middleware.go
│   │   │   └── logging.go
│   │   ├── routes
│   │   │   └── order_routes.go
│   │   ├── models
│   │   │   └── checkout.go
│   ├── app/
│   │   ├── commands/
│   │   │   ├── handler/
│   │   │   │   ├── create_order_handler.go
│   │   │   │   ├── update_order_handler.go
│   │   │   │   ├── pay_order_handler.go
│   │   │   │   ├── cancel_order_handler.go
│   │   │   │   ├── ...
│   │   │   │   └── event_handler.go
│   │   │   ├── command_bus.go
│   │   │   ├── command_handler.go
│   │   │   └── ...
│   │   ├── events
│   │   │   ├── handler/
│   │   │   │   ├── order_created_handler.go
│   │   │   │   ├── order_cancelled_handler.go
│   │   │   │   └── order_cancelled_handler.go
│   │   │   ├── event_bus.go
│   │   │   ├── event_handler.go
│   │   │   └── ...
│   │   ├── queries
│   │   │   ├── get_order_query.go
│   │   │   ├── get_orders_by_user_query.go
│   │   │   ├── get_orders_query.go
│   │   │   └── ...
│   │   ├── app.go
│   ├── domain
│   │   ├── commands
│   │   │   ├── create_order_command.go
│   │   │   ├── cancel_order_command.go
│   │   |   ├── ...
│   │   ├── entities
│   │   │   ├── order.go
│   │   │   ├── order_item.go
│   │   ├── events
│   │   │   ├── events.go
│   │   │   ├── order_created_event.go
│   │   │   ├── order_updated_event.go
│   │   │   ├── order_cancelled_event.go
│   │   │   ├── order_paid_event.go
│   │   │   ├── ...
│   │   ├── repositories
│   │   │   ├── order_repository.go
│   │   │   ├── order_item_repository.go
│   │   │   ├── ...
│   │   ├── services
│   │   │   ├── email_service.go
│   │   │   ├── payment_method.go
│   │   │   ├── ...
│   │   └── value_objects
│   │       ├── order_status.go
│   │       ├── ...
│   │       └── payment_method.go
│   ├── infrastructure/
│   │   ├── configs/
│   │   │   └── config.yaml
│   │   ├── eventstore/
│   │   │   ├── event_store.go
│   │   │   ├── mysql_event_store.go
│   │   │   └── ...
│   │   ├── kafka/
│   │   │   ├── kafka_producer.go
│   │   │   ├── kafka_consumer.go
│   │   │   └── ...
│   │   └── persistence/
│   │       ├── order_repository.go
│   │       ├── ...
│   │       └── order_item_repository.go
│   └── shared/
│       └── errors/
│           ├── app_error.go
│           └── validation_error.go
├── pkg/
│   ├── apperrors/
│   │   ├── apperror.go
│   │   └── ...
│   ├── logging/
│   │   ├── logger.go
│   │   └── ...
│   └── tracing/
│       ├── tracer.go
│       └── ...
├── scripts/
|   └── run_local.sh
├── docs/
│   ├── api/
│   │   └── swagger.yml
├── deployments/
│   ├── kubernetes/
│   │   ├── order-api/
│   │   │   ├── kustomization.yaml
│   │   │   ├── deployment.yaml
│   │   │   └── ...
│   │   └── order-worker/
│   │       ├── kustomization.yaml
│   │       ├── deployment.yaml
│   │       └── ...
```

Giải thích một số phần quan trọng:

* cmd/: Thư mục này chứa file main.go để khởi động ứng dụng và quản lý các command line arguments.
* internal/: Thư mục chính của ứng dụng, bao gồm tất cả các package của ứng dụng.
* **internal/application:** Thư mục `application` chứa các command handlers, event handlers và queries dùng để handle các command, event và query được gửi tới từ các client (web, mobile, ...). Nó cũng chứa file `main.go` để khởi tạo service và các cấu hình khác liên quan đến service.
* **internal/domain/:** Thư mục domain chứa các entities, value objects, repositories, services và các file liên quan đến các command, event và query.
* **internal/app/infrastructure/:** Package chứa các dịch vụ hạ tầng cho ứng dụng, bao gồm cả cơ sở dữ liệu và các client gọi đến các dịch vụ bên ngoài.

### Application layer
    *  internal/app/application/commands/: Package chứa các commands để thay đổi trạng thái của ứng dụng.
    *  internal/app/application/queries/: Package chứa các queries để truy xuất dữ liệu từ ứng dụng.
    *  internal/app/application/services/: Package chứa các services để thực hiện các nhiệm vụ phức tạp trong ứng dụng.
### Domain layer
1. **internal/app/domain/commands/:** Package commands chứa các trạng thái
Command trong Domain được sử dụng để thay đổi trạng thái của các đối tượng trong Domain, trong khi Command trong Application được sử dụng để gửi yêu cầu đến Microservices, Services hoặc các hệ thống khác để thực hiện các tác vụ liên quan đến chức năng của ứng dụng.

Về cơ bản, Command trong Domain chỉ được sử dụng trong phạm vi của Domain của mình, còn Command trong Application được sử dụng để phối hợp các hoạt động của nhiều Domain và hệ thống khác nhau. Command trong Domain chứa thông tin về những thay đổi cần được thực hiện trên đối tượng trong Domain, trong khi Command trong Application chứa thông tin về hành động cần được thực hiện trên toàn bộ ứng dụng hoặc các hệ thống khác.  

Ví dụ, ta cùng xây code sau:
```go
type CreateOrderCommand struct {
    OrderID    string
    CustomerID string
    OrderDate  time.Time
    Total      float64
}

func (c CreateOrderCommand) Validate() error {
    if c.OrderID == "" {
        return fmt.Errorf("order id is required")
    }
    if c.CustomerID == "" {
        return fmt.Errorf("customer id is required")
    }
    if c.Total <= 0 {
        return fmt.Errorf("total must be greater than zero")
    }
    return nil
}

```

2. **internal/app/domain/events/:** Package commands chứa các trạng 

### Infrastructure layer
    *  internal/app/infrastructure/: Package chứa các thành phần cơ sở hạ tầng cho ứng dụng, bao gồm các thành phần lưu trữ, web, và các middlewares.
    