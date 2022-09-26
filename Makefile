gen-proto:
	protoc -I ./protobuf --go_out=. --go-grpc_out=. protobuf/auth_service.proto
	protoc -I ./protobuf --go_out=. --go-grpc_out=. protobuf/user_service.proto