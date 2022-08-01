package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"

	casbin "github.com/casbin/casbin/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	redis "github.com/go-redis/redis/v8"
	"github.com/huavanthong/microservice-golang/email-grpc/gapi"
	pb "github.com/huavanthong/microservice-golang/email-grpc/proto/email"
	"github.com/huavanthong/microservice-golang/user-api-v3/config"
	"github.com/huavanthong/microservice-golang/user-api-v3/controllers"
	"github.com/huavanthong/microservice-golang/user-api-v3/routes"
	"github.com/huavanthong/microservice-golang/user-api-v3/services"
	"github.com/huavanthong/microservice-golang/user-api-v3/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	docs "github.com/huavanthong/microservice-golang/user-api-v3/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	// Server setting
	server      *gin.Engine     // The framework's instance, it contains the muxer, middleware and configuration settings.
	ctx         context.Context // Context running in background
	mongoclient *mongo.Client   // MongoDB
	redisclient *redis.Client   // For in-memory data store

	// User Controller setting
	userService         services.UserService
	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	// Admin Controller setting
	adminService         services.AdminService
	AdminController      controllers.AdminController
	AdminRouteController routes.AdminRouteController

	// Authenticate Controller setting
	authCollection         *mongo.Collection
	authService            services.AuthService
	AuthController         controllers.AuthController
	AuthRouteController    routes.AuthRouteController
	SessionRouteController routes.SessionRouteController

	// 👇 Add the Post Service, Controllers and Routes
	postService         services.PostService
	PostController      controllers.PostController
	postCollection      *mongo.Collection
	PostRouteController routes.PostRouteController
)

func init() {

	// Loading config from variable environment
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	// Init an context running in background
	ctx = context.TODO()

	// Connect to MongoDB
	mongoconn := options.Client().ApplyURI(config.DBUri)
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		panic(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("MongoDB successfully connected...")

	// Connect to Redis
	redisclient = redis.NewClient(&redis.Options{
		Addr: config.RedisUri,
	})

	if _, err := redisclient.Ping(ctx).Result(); err != nil {
		panic(err)
	}

	err = redisclient.Set(ctx, "test", "Welcome to Golang with Redis and MongoDB", 0).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Redis client connected successfully...")

	// Init Collections
	authCollection = mongoclient.Database("golang_mongodb").Collection("users")
	// Setting service
	userService = services.NewUserServiceImpl(authCollection, ctx)
	adminService = services.NewAdminServiceImpl(authCollection, ctx)
	authService = services.NewAuthService(authCollection, ctx)
	// Setting controller
	AuthController = controllers.NewAuthController(authService, userService, ctx, authCollection)
	AuthRouteController = routes.NewAuthRouteController(AuthController)
	SessionRouteController = routes.NewSessionRouteController(AuthController)

	AdminController = controllers.NewAdminController(adminService)
	AdminRouteController = routes.NewRouteAdminController(AdminController)

	UserController = controllers.NewUserController(userService)
	UserRouteController = routes.NewRouteUserController(UserController)

	// 👇 Add the Post Service, Controllers and Routes
	postCollection = mongoclient.Database("golang_mongodb").Collection("posts")
	postService = services.NewPostService(postCollection, ctx)
	PostController = controllers.NewPostController(postService)
	PostRouteController = routes.NewPostControllerRoute(PostController)

	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	server = gin.Default()
}

// @title UserManagement Service API Document
// @version 1.0
// @description List APIs of UserManagement Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8000
// @BasePath /api/v3
func main() {

	/************************ Init MongoDB *************************/
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load config", err)
	}

	defer mongoclient.Disconnect(ctx)

	/************************ Start internal server *************************/
	startGinServer(config)
	startGrpcServer(config)

}

func startGrpcServer(config config.Config) {
	// create an instance of the Email server
	server, err := gapi.NewGrpcServer(config, authService, userService, authCollection)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	// create a new gRPC server, use WithInsecure to allow http connections
	grpcServer := grpc.NewServer()

	// register the email server
	pb.RegisterAuthServiceServer(grpcServer, server)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcServer)

	// accept connection for grcp port
	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("Unable to create listener: %s, err: %s", config.GrpcServerAddress, err)
	}

	// listen for requests
	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}

func startGinServer(config config.Config) {

	/************************ Connect Redis *************************/
	value, err := redisclient.Get(ctx, "test").Result()
	if err == redis.Nil {
		fmt.Println("[Main] key: test does not exist")
	} else if err != nil {
		panic(err)
	}

	/************************ Allow Cross Orgin Resource Sharing  *************************/
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{config.Origin}
	// corsConfig.AllowCredentials = true

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	docs.SwaggerInfo.BasePath = "/api/v3"

	/************************ Init GIN session  *************************/
	// generate google token
	token, err := utils.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
	})

	server.Use(sessions.Sessions("goquestsession", store))

	/************************ Server routing  *************************/
	router := server.Group("/api/v3")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": value})
	})

	/************************ Policy to authorize role user  *************************/
	// load the casbin model and policy from files, database is also supported.
	authorizeMiddeware, _ := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")

	/************************ Controller  *************************/
	AuthRouteController.AuthRoute(router, userService, authorizeMiddeware)
	UserRouteController.UserRoute(router, userService)
	AdminRouteController.AdminRoute(router, adminService)
	SessionRouteController.SessionRoute(router)
	// 👇 Evoke the PostRoute
	PostRouteController.PostRoute(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Fatal(server.Run(":" + config.Port))
}
