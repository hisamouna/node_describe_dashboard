package main

import (
	"log"
	"net"

	"github.com/hisamouna/node_describe_dashboard/handler"
	"github.com/hisamouna/node_describe_dashboard/pkg/server/node"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen. : %v", err)
	}

	s := grpc.NewServer()
	node.RegisterNodeServiceServer(s, &handler.NodeHandler{})
	log.Printf("Serving gRPC on 0.0.0.0:%s", port)
	log.Fatal(s.Serve(lis))
}
