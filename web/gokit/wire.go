package gokit

import (
	"MiniDNS2/service"
	"MiniDNS2/web/gokit/endpoint"
	"MiniDNS2/web/gokit/transport"
)

type server struct {
	port       string
	grpcServer *transport.KitGRPCServer
}

func NewServer(port string) (*server, error) {
	//连接Mysql
	//连接redis
	//新建dao.Dao
	endpoints := endpoint.NewEndpoints(service.Srvc)
	DNSServer := transport.NewKitGRPCServer(endpoints)
	ret := &server{
		port:       port,
		grpcServer: DNSServer,
	}
	return ret, nil
}

/*
func (s *server)Run() error {
	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	svr := grpc.NewServer(opts)
	mux := http.NewServeMux()
	http.Serve(lis, mux)
}*/
