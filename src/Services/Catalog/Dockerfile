## We specify the base image we need for our go application
FROM golang:1.17-alpine

## Create an /product-api directory within our image
## that will hold our application source files
RUN mkdir /catalog-api

## Copy everything in the root directory into our /catalog-api directory
ADD . /catalog-api

## Specify /product-api directory for execution
WORKDIR /catalog-api

# Build binary file
RUN go build -o catalog-api cmd/catalog-api/main.go

# chạy container
CMD ["./catalog-api"]
