Init Project
```azure
go mod tidy
```

Run Controller

```azure
cd pkg/controller
go run main.go
```

Run Agent
```azure
cd pkg/agent
go run main.go
```

Generate Protobuf Go Files
```azure
protoc --go_out=. --go-grpc_out=. pkg/common/pb/PROTO_FILE_NAME
```

Noauth Mode
``` 
environment variables
// Noauth enabled 
NO_AUTH="TRUE"
// Noauth disabled
NO_AUTH="FALSE"
```

Run Mode
``` 
"PRODUCTION"/"DEVELOP"
```