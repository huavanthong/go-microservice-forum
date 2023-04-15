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
  * internal/app/: Package chứa tất cả các thành phần của ứng dụng, bao gồm các commands, queries, services, domain entities, repositories, value objects và các thành phần chia sẻ.
    *  internal/app/application/commands/: Package chứa các commands để thay đổi trạng thái của ứng dụng.
    *  internal/app/application/queries/: Package chứa các queries để truy xuất dữ liệu từ ứng dụng.
    *  internal/app/application/services/: Package chứa các services để thực hiện các nhiệm vụ phức tạp trong ứng dụng.
    *  internal/app/domain/: Package chứa các entities, events, repositories và value objects cho domain của ứng dụng.
    *  internal/app/infrastructure/: Package chứa các thành phần cơ sở hạ tầng cho ứng dụng, bao gồm các thành phần lưu trữ, web, và các middlewares.
    *  internal/app/shared/: Package chứa các thành phần chia sẻ trong ứng dụng, bao gồm các lỗi
