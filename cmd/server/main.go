package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	rpc "github.com/devbv/go-server-boilerplate/pkg/rpc"
	"github.com/devbv/go-server-boilerplate/pkg/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server := server.New()
	mux := runtime.NewServeMux()
	err := rpc.RegisterTODOServiceHandlerServer(ctx, mux, server)
	if err != nil {
		return
	}

	// Start HTTP server
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		return
	}
}
