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
// auth enabled 
AUTH_ENABLED="TRUE"
// auth disabled
AUTH_ENABLED="FALSE"
```

Run Mode
``` 
"PRODUCTION"/"DEVELOP"
```

Go Test
```azure
cd pkg/k8s_test
go test
--  for specific func
go test -run TestFunctionA
-- coverage
go tool cover --html=coverage.out
-- to view in browser 
go tool cover --html=coverage.out
```