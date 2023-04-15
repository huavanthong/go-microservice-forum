# Introduction


# Project Structure
```lua 
-ordering-microservices/
  -cmd/
    -ordering-api/
      -main.go
  -internal/
    -app/
      -application/
        -commands/
          -create_order_command.go
          -update_order_command.go
          -cancel_order_command.go
          -pay_order_command.go
        -queries/
          -get_order_query.go
          -get_orders_by_user_query.go
          -get_orders_query.go
        -app.go
      -domain/
        -entities/
          -order.go
        -events/
          -event.go
          -order_created_event.go
          -order_updated_event.go
          -order_cancelled_event.go
          -order_paid_event.go
        -repositories/
          -order_repository.go
          -event_store_repository.go
        -valueobjects/
          -order_item.go
          -money.go
      -infrastructure/
        -persistence/
          -db/
            -migrations/
              -20220413_001_init.sql
            -postgres/
              -postgres.go
          -eventstore/
            -event_store.go
        -web/
          -controllers/
            -order_controller.go
          -middlewares/
            -auth_middleware.go
            -cors_middleware.go
      -shared/
        -errors/
          -app_error.go
          -validation_error.go
    -pkg/
      -config/
        -config.go
      -logger/
        -logger.go
  -configs/
    -config.yaml
  -docs/
    -api/
      -swagger.yaml
  -scripts/
    -run_local.sh
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
