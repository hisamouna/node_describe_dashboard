package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hisamouna/node_describe_dashboard/handler"
	"github.com/hisamouna/node_describe_dashboard/pkg/server/node"
	"google.golang.org/grpc"
)

const (
	target     = "0.0.0.0"
	targetPort = ":8080"
	port       = ":8090"
)

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s%s", target, targetPort),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	gwmux := runtime.NewServeMux()
	mux := http.NewServeMux()

	err = node.RegisterNodeServiceHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}
	mux.HandleFunc("/", gwmux.ServeHTTP)

	/*
		gwServer := &http.Server{
			Addr:    port,
			Handler: gwmux,
		}
		log.Printf("Serving gRPC-Gateway on http://0.0.0.0%s", port)
		log.Fatal(gwServer.ListenAndServe())
	*/

	log.Printf("Serving gRPC-Gateway on http://0.0.0.0%s", port)
	log.Fatal(http.ListenAndServe(port, handler.AllowCORS(mux)))
}
