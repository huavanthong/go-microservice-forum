.PHONY: proto

proto:
	protoc --proto_path=proto ./proto/*.proto --go_out=  --go-grpc_out=. 

prototemp:
	protoc --proto_path=proto ./proto/email.proto --go_out=. --go-grpc_out=.


# --go_opt=paths=source_relative \
	 --go-grpc_opt=paths=source_relative \