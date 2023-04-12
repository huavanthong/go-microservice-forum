FROM golang:1.17-alpine AS base
WORKDIR /app
EXPOSE 80


## We specify the base image we need for our go application
FROM golang:1.17-alpine AS build
## Create an directory within our image
## that will hold our application source files
WORKDIR /src

# Copy the go.mod and go.sum files to the working directory
COPY ["Services/Discount/go.mod", "Services/Discount/"]
COPY ["Services/Discount/go.sum", "Services/Discount/"]

## Copy everything in the root directory into our directory
COPY . .

## Specify directory for execution
WORKDIR "/src/Services/Discount/"

# Build binary file
RUN go build -o discount-api ./cmd/discount-api/main.go

EXPOSE 8002

# Run command
CMD ["./discount-api"]
