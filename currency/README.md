# Introduction
The currency service is a gRPC service which provides up to date exchange rates and currency conversion capabilities.

1. To install protoc for server. [here](#install-protos)
2. To build out *.pb.go. [here](#build-protos)
3. 

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

