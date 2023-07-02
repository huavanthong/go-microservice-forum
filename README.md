# go-microservice-forum


## Reference
## Design with Golang
* How to design
    - Tutorial for designing on Golang. Refer: [here](https://github.com/techschool/simplebank)
    - Design Pattern with experience. Refer: [here](https://github.com/iluwatar/java-design-patterns)
    - Design RESTful API with microservice, perfect design. Refer: [here](https://github.com/restuwahyu13/gin-rest-api)
* Roadmap
    - For developer with Golang. Refer: [here](https://github.com/Alikhll/golang-developer-roadmap)
* Database
    - Think about database. Refer: [here](https://dev.to/fabric_commerce/building-a-hybrid-sql-nosql-e-commerce-data-model-3fc3)
## Architecture
* Overview
![110304529-c5b70180-800c-11eb-832b-a2751b5bda76](https://user-images.githubusercontent.com/50081052/188884348-e9c10440-77d8-4edc-a578-b0b158df43e0.png)
* Project
    - Project ASP.NET example for demo microservice. Refer: [here](https://mehmetozkaya.medium.com/aspnetrun-microservices-renewed-d08901b5e06f), [doc](https://medium.com/aspnetrun/microservices-architecture-on-net-3b4865eea03f), [source](https://github.com/aspnetrun/run-aspnetcore-microservices)

## Microservice
### Product microservice
* Logger
    - Design logger for Golang project - product service. [Refer](https://techmaster.vn/posts/36655/su-dung-uber-zap-thay-the-cho-logging-mac-dinh-cua-golang)

### gRPC (g Remote Procedure Call)
gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.
* Purpose
    - When we need to use gRPC. Refer: [here](https://www.wallarm.com/what/the-concept-of-grpc)
* gRPC
    - Communication between microservice. Refer: [here](https://techdozo.dev/grpc-for-microservices-communication/)
    - Traefik with gRPC. Refer: [here](https://doc.traefik.io/traefik/user-guides/grpc/)
    - What is Traefik. Refer: [here](https://www.devopsschool.com/blog/what-is-traefik-how-to-learn-traefik/)

### Basket microsrevice
Developing Basket microservice which includes:
* REST API principles, CRUD operations.
* Redis DB NoSQL database connection and containerization
* Mongo DB database for storing customer's basket.
* Implement Basket microservice using the principles of Clean Architecture, which emphasizes separation of concerns and modularization for creating maintainable and scalable software.
* Swagger Open API implementation.
* Dockerfile implementation.
### Email microservice
* Project
    - Example project email using **RabbitMQ**. Refer: [here](https://github.com/savsgio/microservice-email)
    - Document for email microserivce. Refer: [here](https://www.cloudbees.com/blog/email-as-a-microservice)

### Payment microservice
* Project
    - Example project email using **RabbitMQ**. Refer: [here](https://github.com/savsgio/microservice-email)
    - Document for email microserivce. Refer: [here](https://www.cloudbees.com/blog/email-as-a-microservice)
### User microservice
* Project
    - Example project user-microservice. Refer: [here](https://github.com/raycad/go-microservices)
    - Architecture for user-microserivce. Refer: [here](https://github.com/huavanthong/microservice-golang/tree/master/reference/user-api#architecture)
    - Perfect design for user-microservice, Refer: [here](https://github.com/wpcodevo/golang-mongodb-api/tree/golang-mongodb-reset-password)
* Validator
    - Implement validator using go-playground. Refer: [here](https://github.com/go-playground/validator)
    - Implement validator using go-playground through GIN framework. Refer: [here](https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/)
* Authentication
    - Implement Google OAuth2 with GIN. Refer: [here](https://skarlso.github.io/2016/06/12/google-signin-with-go/)
    - Implement JWT for login account. Refer: [here](https://tienbm90.medium.com/authentication-and-authorization-in-gin-application-with-jwt-and-casbin-a56bbbdec90b)
    - Implement login/logout feature using session in gin. Refer: [here](https://github.com/Depado/gin-auth-example)
    - Implement a multi-level authentication system for goalng. [here](https://mattermost.com/blog/how-to-build-an-authentication-microservice-in-golang-from-scratch/)
    - Implement feature to change password. [here](https://auth0.com/docs/authenticate/database-connections/password-change)
* Authorization
    - Implement authorize role. Refer: [here](https://www.zupzup.org/casbin-http-role-auth/)
    - Implement authorize role with GIN framework. Refer: [here](https://github.com/gin-contrib/authz)
    - User Management Roles and Functions. Refer: [here](https://www.ibm.com/docs/en/strategicsm/10.1.1?topic=roles-user-management-functions)
* Security
    - Convert user's password in plain text to bcrypt. [here](https://github.com/huavanthong/microservice-golang/blob/master/reference/user-api/security/bcrypt.go)
    - Avoid SQL Injection. Refer: [here](https://github.com/huavanthong/microservice-golang/blob/master/reference/user-api/security/bcrypt.go)
* Database.
    - MongoDB
        - Design database for user-microservice. Refer: [here](https://github.com/huavanthong/microservice-golang/tree/master/user-api#design-database-for-user-microservice)   
    - PostgreSQL
        - Migrate database to PostgreSQL. Refer: [here](https://dev.to/techschoolguru/how-to-write-run-database-migration-in-golang-5h6g)

### Payment microservice

## User microservice version 3

* Microservice pattern
    - Circuit pattern. Refer: [here](https://medium.com/nerd-for-tech/design-patterns-for-microservices-circuit-breaker-pattern-ba402a45aac2)

## Docker
1. Connect go-app with mongo db manually. Refer: [here](https://wahyu-ehs.medium.com/golang-and-mongodb-in-docker-cab7773c04ce)
2. Build out image Docker for Go program, and knowledge for build CI/CD. Refer to private info at: [here](https://www.meisternote.com/app/note/G1vMneveYIrP/build-image-docker-for-go)
## Github Action
1. Apply CI/CD for project. Refer: [here](https://docs.github.com/en/actions/quickstart)
