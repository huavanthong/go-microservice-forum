.PHONY: proto

proto:
	protoc --proto_path=proto --go_out=./proto/email --go_opt=paths=source_relative \
							--go-grpc_out=./proto/email --go-grpc_opt=paths=source_relative proto/*.proto

