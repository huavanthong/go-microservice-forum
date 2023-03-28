## We specify the base image we need for our go application
FROM golang:1.17-alpine

## Create an directory within our image
## that will hold our application source files
RUN mkdir /basket-api

## Copy everything in the root directory into our directory
ADD . /basket-api

## Specify directory for execution
WORKDIR /basket-api

# Build binary file
RUN go build -o basket-api cmd/basket-api/main.go

# chạy container
CMD ["./basket-api"]
