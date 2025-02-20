# Development Guide

We welcome contributions! If you'd like to contribute to Lighthouse, please:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request.

## Prerequisites
- A Kubernetes Cluster (Can run locally with Kind of other tools) \
or a `kubeconfig` file of a remotely accessible Kubernetes cluster
- A mongodb instance (For now, we only supporting mongodb but have plans to support more)
- Golang dev environment

## Steps
#### 1. Create  a Kind Cluster (Optional)
```azure
kind create cluster --name lighthouse
```

#### 2. Clone Repository
#### 3. Create a `.env` file taking reference from `.env-example` file
#### 4. Set Environment Variables (in .env file):
- Set `RUN_MODE=DEVELOP`
- Set `KUBE_CONFIG_FILE=YOUR_KUBECONFIG_FILENAME`. Kubeconfig file is expected in `$HOME/.kube` directory.
- Set `IS_INTERNAL_SERVER="TRUE"`
- Set `MONGO_URI=YOUR_MONGO_CONNECTION_URI`
#### 5. Download Dependencies
```azure
go mod tidy
```

#### 6. Run Controller

```azure
go run pkg/controller/main.go
```

#### 7. Run Agent
```azure
go run pkg/agent/main.go
```

#### 7. Run Frontend
```azure
cd frontend
npm install
npm start
```

## Others
### 1. Generate Protobuf Go Files
```azure
protoc --go_out=. --go-grpc_out=. pkg/common/pb/PROTO_FILE_NAME
```

### 2. Run Go Test
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