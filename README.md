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

└─> go run cmd/gateway/main.go
2021/09/13 23:22:15 Serving gRPC-Gateway on http://0.0.0.0:8090

└─> curl http://localhost:8090/v1/node/describe | jq -r '.'

{
  "nodes": [
    {
      "name": "kinda-control-plane",
      "pods": [
        {
          "namespace": "kube-system",
          "name": "etcd-kinda-control-plane"
        },
        {
          "namespace": "kube-system",
          "name": "kindnet-c8fhf"
        },
.
.
.
```
