Init Project
```azure
go mod tidy
```

Run Controller

```azure
cd pkg/controller
go run main.go
```

Run Worker
```azure
cd pkg/worker
go run main.go
```

Generate Protobuf Go Files
```azure
protoc --go_out=. --go-grpc_out=. pkg/common/pb/PROTO_FILE_NAME
```