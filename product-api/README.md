# Introduction

# Getting Started
* You can run product-api on docker.
* You can run product-api at local host.
### Run on docker
To build docker image for product-api. 
```
docker build -t product-api .
```

To run product-api image on docker. 
```
docker run -p 8080:8080 -it product-api
```
### Run local host
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

To run API on server
```
curl -X GET localhost:8080/ 
curl localhost:8080/ -X POST -d '{"Name": "New Product", "Price": 1.23, "SKU":"abc-def-ghi"}'
```