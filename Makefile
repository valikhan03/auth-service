gen-proto:
	protoc -I ./protobuf --go_out=. --go-grpc_out=. protobuf/auth_service.proto
