# Introduction
The currency service is a gRPC service which provides up to date exchange rates and currency conversion capabilities.

# Table of Contents
1. To install protoc for server. [here](#install-protos)
2. To build out *.pb.go. [here](#build-protos)
3. To run server and use it. [here](#getting-started)
4. To test package data. [here](#data-package)
5. To get coverage result. [here](#for-coverage-testing)
6. To get benchmark result. [here](#for-benchmarks-testing)


## Install protos
To build the gRPC client and server interfaces, first install protoc:
### Linux
```shell
sudo apt install protobuf-compiler
```

### Mac
```shell
brew install protoc
```

### Windows
Download protoc win64 zip package. Extract it and put to GOPATH/bin.
Link: [protoc win64](https://github.com/protocolbuffers/protobuf/releases)
```
    protoc-3.20.1-win64.zip
```
More details: [here](https://www.youtube.com/watch?v=ES_GI-lmhEU)

## Build protos
Old command for building
```shell
protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
```

Update command for building
```shell
protoc --proto_path=proto ./proto/currency.proto --go_out=. --go-grpc_out=.
```

**Note:**
- go_out: will build out file *.pb.go
- go-grpc_out: will build out file *_grpc.pb.go
- Please understand the difference between them

## Getting Started
### Install grpccurl
To test the system install `grpccurl` which is a command line tool which can interact with gRPC API's
```
go get github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
More details: [here](https://github.com/fullstorydev/grpcurl)

### List Services
```
grpcurl --plaintext localhost:9092 list
Currency
grpc.reflection.v1alpha.ServerReflection
```

### List Methods
```
grpcurl --plaintext localhost:9092 list Currency        
Currency.GetRate
```

### Method detail for GetRate
```
grpcurl --plaintext localhost:9092 describe Currency.GetRate
Currency.GetRate is a method:
rpc GetRate ( .RateRequest ) returns ( .RateResponse );
```

### RateRequest detail
```
grpcurl --plaintext localhost:9092 describe .RateRequest    
RateRequest is a message:
message RateRequest {
  string Base = 1 [json_name = "base"];
  string Destination = 2 [json_name = "destination"];
}
```

### Execute a request
```
grpcurl --plaintext -d '{"base": "GBP", "destination": "USD"}' localhost:9092 Currency/GetRate
{
  "rate": 0.5
}
```
### Execute a request for SubscribeRates

The parameter `-d @` means that gRPCurl will read the messages from StdIn.

```
grpcurl --plaintext --msg-template -d @ localhost:9092 Currency/SubscribeRates 
```

You can send a message to the server using the following payload

```
{
  "Base": "EUR",
  "Destination": "EUR"
}
```

## Testing with Evans
As you can see, the above demonstration help you understand and know how to test gRPC with command grpcurl.  
However, we also have a powerful tool to help you test with gRPC. It is Evans
### To install it
Refer: 
* https://github.com/ktr0731/evans#installation
### Started gRCP with Evans
**Step 1:** Start your currency service.
```
./currency.exe
```
**Step 2:**  Start Evans tool with your gRPC server.
```
evans --host localhost --port 9092 -r repl
```
## Data package
Reference: 
* [guideline for testing](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
### Entering to place want to test
To want to test data on currency package
```
cd currency/data
```
More details: [here]()
### For testing
For running the testing data
```
go test
```

For a additional detail info
```
go test -v
```
### For coverage testing
To get a result of coverage in data package
```
go test -cover
```

#### Windows
To get coverage result analytics
```
go test -coverprofile=coverage
```

To get html template from coverage result
```
go tool cover -html=coverage
```
#### Linux
To get coverage result analytics
```
go test -coverprofile=coverage.out
```

To get html template from coverage result
```
go tool cover -html=coverage.out
```
### For benchmarks testing
Run testing for benckmark
```
go test -bench=.
```

You can also declare benchmark functions explicitly:
```
go test -bench=NewRates
go test -bench=GetRate
```