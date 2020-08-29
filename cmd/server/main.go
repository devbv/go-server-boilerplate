package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/devbv/go-server-boilerplate/internal/server"
	"github.com/devbv/go-server-boilerplate/pkg/todo"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server := server.New()
	mux := runtime.NewServeMux()
	err := todo.RegisterTODOServiceHandlerServer(ctx, mux, server)
	if err != nil {
		return
	}

	// Start HTTP server
	err = http.ListenAndServe(":8000", mux)
	if err != nil {
		return
	}
}
