FROM golang:1.17-alpine AS base
WORKDIR /app
EXPOSE 80


## We specify the base image we need for our go application
FROM golang:1.17-alpine AS build
## Create an directory within our image
## that will hold our application source files
WORKDIR /src

# Install curl
RUN apk update && apk add --no-cache curl

# Copy the go.mod and go.sum files to the working directory
COPY ["Services/Basket/go.mod", "Services/Basket/"]
COPY ["Services/Basket/go.sum", "Services/Basket/"]

## Copy everything in the root directory into our directory
COPY . .

## Specify directory for execution
WORKDIR "/src/Services/Basket/"

# Chỉ định cổng mà container sẽ lắng nghe các kết nối đến. 
EXPOSE 8001

# Build binary file
RUN go build -o basket-api ./cmd/basket-api/main.go

# Run command
CMD ["./basket-api"]
