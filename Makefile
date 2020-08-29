
install:
	go mod tidy
	go install \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
    github.com/golang/protobuf/protoc-gen-go

generate:
	protoc --proto_path=./rpc --grpc-gateway_out=./pkg/todo --swagger_out=./rpc --go_out=plugins=grpc:./pkg/todo rpc/*.proto