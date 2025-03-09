gen:
	 @protoc \
	 --proto_path=pkg/common/pb pkg/common/pb/controller.proto \
	 --go_out=./pkg/common/pb \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=./pkg/common/pb \
	 --go-grpc_opt=paths=source_relative