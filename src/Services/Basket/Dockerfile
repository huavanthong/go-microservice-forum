# Base stage
FROM golang:1.20-buster AS base
WORKDIR /app
EXPOSE 80

# Build stage
FROM golang:1.20-buster AS build
## Create an directory within our image
## that will hold our application source files
WORKDIR /src

# Install curl
RUN apt-get update && apt-get install -y curl

# Copy the go.mod and go.sum files to the working directory
COPY ["Services/Basket/go.mod", "Services/Basket/"]
COPY ["Services/Basket/go.sum", "Services/Basket/"]

## Copy everything in the root directory into our directory
COPY . .

## Specify directory for execution
WORKDIR "/src/Services/Basket/"

# Specify the port that container will listen all coming connection.
EXPOSE 8001

# Build binary file
RUN go build -o basket-api ./cmd/basket-api/main.go

# Run command
CMD ["./basket-api"]
