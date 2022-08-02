module github.com/huavanthong/microservice-golang/product-api-v3

go 1.14

require (
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.8.1
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/viper v1.12.0
	github.com/stretchr/testify v1.8.0
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.1
	go.mongodb.org/mongo-driver v1.8.3
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/huavanthong/microservice-golang/currency => ../currency
