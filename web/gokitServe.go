package web

import (
	"MiniDNS2/library"
	"MiniDNS2/proto"
	"MiniDNS2/service"
	"MiniDNS2/web/gokit/endpoint"
	"MiniDNS2/web/gokit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func GoKitServe(port string) {
	goKitServeMux := http.NewServeMux()
	endpoints := endpoint.NewEndpoints(service.Srvc)
	var goKitServer *kithttp.Server
	goKitServer = kithttp.NewServer(endpoints.GetIPEndpoint, transport.GetIPDecodeRequest, transport.GetIPEncodeResponse)
	goKitServeMux.Handle("/getip", goKitServer)
	goKitServer = kithttp.NewServer(endpoints.InsertEndpoint, transport.InsertDecodeRequest, transport.InsertEncodeResponse)
	goKitServeMux.Handle("/insert", goKitServer)
	goKitServer = kithttp.NewServer(endpoints.UpdateEndpoint, transport.UpdateDecodeRequest, transport.UpdateEncodeResponse)
	goKitServeMux.Handle("/update", goKitServer)
	goKitServer = kithttp.NewServer(endpoints.DeleteEndpoint, transport.DeleteDecodeRequest, transport.DeleteEncodeResponse)
	goKitServeMux.Handle("/delete", goKitServer)
	err := http.ListenAndServe(port, goKitServeMux)
	library.Check(err, "http.ListenAndServe error in web.GoKitServe")
}

func GoKitGRPCServe(port string) {
	endpoints := endpoint.NewEndpoints(service.Srvc)
	goKitGRPCServer := transport.NewKitGRPCServer(endpoints)
	lis, err := net.Listen("tcp", port)
	library.Check(err, "net.Listen error in web.gokitServe.GoKitGRPCServe")
	grpcServer := grpc.NewServer()
	proto.RegisterDNSServer(grpcServer, goKitGRPCServer)
	err = grpcServer.Serve(lis)
	library.Check(err, "grpcServer.Serve(lis) error in web.gokitServe.GoKitGRPCServe")
}
