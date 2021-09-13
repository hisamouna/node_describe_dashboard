# node_describe_dashboard

## Structure

```sh
.
├── backend
│   ├── cmd
│       ├── gateway : gRPC-Gateway
│       │   └── main.go
│       └── server
│           └── main.go : gRPC server
└── frontend
```

## Buckend

### Execute
```sh
└─> go run cmd/server/main.go
2021/09/13 23:22:10 Serving gRPC on 0.0.0.0::8080

└─> go run cmd/server/main.go
2021/09/13 23:22:10 Serving gRPC on 0.0.0.0::8080
```
