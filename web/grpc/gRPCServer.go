package grpc

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"MiniDNS2/proto"
	"MiniDNS2/service"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"net"
)

func GRPCServe(port string) {
	lis, err := net.Listen("tcp", port)
	library.Check(err, "net.listen err in web.GRPCServe")
	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(service.UnaryServerInterceptor)))
	proto.RegisterDNSServer(s, &GRPCServer{})
	err = s.Serve(lis)
	library.Check(err, "grpc.server.serve error in web.GRPCServe")
}

type GRPCServer struct{}

func (grpc *GRPCServer) GetIP(ctx context.Context, r *proto.GetReq) (*proto.GetResp, error) {
	domain := r.Domain
	req := &model.GetReq{Domain: domain} //service
	resp := service.Srvc.GetIP(ctx, req) //service
	ret := &proto.GetResp{Domain: resp.Domain, IPs: resp.IPs}
	return ret, *new(error)
}

func (grpc *GRPCServer) Insert(ctx context.Context, r *proto.InsertReq) (*proto.InsertResp, error) {
	domain := r.Domain
	ip := r.IP
	req := &model.InsertReq{Domain: domain, IP: ip} //service
	resp := service.Srvc.Insert(ctx, req)           //service
	ret := &proto.InsertResp{Domain: resp.Domain, IP: resp.IP, Result: resp.Result}
	return ret, *new(error)
}

func (grpc *GRPCServer) Update(ctx context.Context, r *proto.UpdateReq) (*proto.UpdateResp, error) {
	dmsrc := r.Domainsrc
	ipsrc := r.IPsrc
	dmdst := r.Domaindst
	ipdst := r.IPdst
	req := &model.UpdateReq{Domainsrc: dmsrc, IPsrc: ipsrc, Domaindst: dmdst, IPdst: ipdst} //service
	resp := service.Srvc.Update(ctx, req)                                                   //service
	ret := &proto.UpdateResp{Affected: int64(resp.Affected), Result: resp.Result}
	return ret, *new(error)
}

func (grpc *GRPCServer) Delete(ctx context.Context, r *proto.DeleteReq) (*proto.DeleteResp, error) {
	domain := r.Domain
	ip := r.IP
	req := &model.DeleteReq{Domain: domain, IP: ip}
	resp := service.Srvc.Delete(ctx, req)
	ret := &proto.DeleteResp{Affected: int64(resp.Affected), Result: resp.Result}
	return ret, *new(error)
}
