package web

import (
	"MiniDNS2/library"
	"MiniDNS2/proto"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

func GatewayServe(port, Endpoint string) {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	err := proto.RegisterDNSHandlerFromEndpoint(ctx, mux, Endpoint, opts)
	library.Check(err, "proto.RegisterDNSHandlerFromEndpoint error in web.GatewayServe")
	err = http.ListenAndServe(port, mux)
	library.Check(err, "http.ListenAndServe error in web.GatewayServe")
}
