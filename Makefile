generate_grpc_error:
	protoc --go-grpc_out=pkg/grpc/error --go_out=pkg/grpc/error pkg/grpc/error/error.proto