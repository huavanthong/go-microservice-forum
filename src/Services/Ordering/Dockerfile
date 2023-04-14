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
COPY ["Services/Ordering/go.mod", "Services/Ordering/"]
COPY ["Services/Ordering/go.sum", "Services/Ordering/"]

## Copy everything in the root directory into our directory
COPY . .

## Specify directory for execution
WORKDIR "/src/Services/Ordering/"

# Chỉ định cổng mà container sẽ lắng nghe các kết nối đến. 
EXPOSE 8004

# Build binary file
RUN go build -o ordering-api ./cmd/ordering-api/main.go

# Run command
CMD ["./ordering-api"]
