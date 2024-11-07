FROM golang:1.17-alpine AS base
WORKDIR /app
EXPOSE 80

## We specify the base image we need for our go application
FROM golang:1.17-alpine AS build

## Create an /product-api directory within our image
## that will hold our application source files
WORKDIR /src

COPY ["Services/Catalog/go.mod", "Services/Catalog/"]
COPY ["Services/Catalog/go.sum", "Services/Catalog/"]

## Copy everything in the root directory into our /catalog-api directory
COPY . .

## Specify /product-api directory for execution
WORKDIR "/src/Services/Catalog/"

# Build binary file
RUN go build -o catalog-api ./cmd/catalog-api/main.go

# chạy container
CMD ["./catalog-api"]
