# Introduction
This project will help you implement a microservice project.

# Question
* [What is the redox middleware?]()
# Table of Contents
* [How to implement a validate for product-api]

# Design
### Desin validate in the same of product 
to update a simple validate for a product.  
* Step 1: create a receiver for product 
* Step 2: use validator package, and create a object.
```go
func (p *Product) Validate() error {
	validate := validator.New()

	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}
```
### Design validate in another directory.
 
# Getting Started
* Run product-api on docker -> [here](#run-on-docker)
* Run product-api on localhost -> [here](#run-local-host)
* Execute API on this server -> [here](#execute-api)
* Run swagger documentation -> [here](#swagger)


## Run on docker
To build docker image for product-api. 
```
docker build -t product-api .
```

To run product-api image on docker. 
```
---port: docker port 8080 -> localhost port: 8080
docker run -p 8080:8080 -it product-api
```
## Run local host
To build on local host
```
product-api> go build
```

To run product-api on local host
```
product-api> go run main
```
## Execute API
To run API on server
```
curl localhost:8080
```

To run API on server with format json
```
curl localhost:8080/ | jq
```
#### GET method
To get all products on server
```
curl -X GET localhost:8080/products
```

To get a specific product on server
```
curl -X GET localhost:8080/products/1
```
#### PUT method
To update a product by specific Id
```
curl localhost:8080/products/1 -X PUT -d '{"name": "Tea", "price": 1.0, "sku":"aada-sdd-ddf"}'
```
#### POST method
To post a new product to server
```
curl -X POST localhost:8080/products -d '{"name": "New Product", "Price": 1.0, "sku":"aada-sdd-ddf"}'
```
##### For testing POST method
```
curl localhost:8080/products -X POST -d '{"name": "Tea", "Price": 1.0, "sku":"aada-sdd"}'
curl localhost:8080/products -X POST -d '{"name": "Water"}'
```
#### DELETE method
```
curl -X DELETE localhost:8080/products/1
```
## Swagger
To generate swagger.yaml file
```
make swagger
```
### Redoc middleware
```
The redox middleware allows us to be able to host our documentation website within our API service.
```

To run swagger.yaml on server with redoc, please code as below
```go
	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

    getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
```

