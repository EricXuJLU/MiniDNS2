package web

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"MiniDNS2/proto"
	"MiniDNS2/service"
	"context"
	"google.golang.org/grpc"
	"net"
)

func GRPCServe(port string) {
	//fout, err := os.Create("test.txt")
	//library.Check(err, "file open failed")
	//defer fout.Close()
	//fout.WriteString("grpc start 1")
	lis, err := net.Listen("tcp", port)
	library.Check(err, "net.listen err in web.GRPCServe")
	//fout.WriteString("grpc start 2")
	s := grpc.NewServer()
	proto.RegisterDNSServer(s, &gRPCserver{})
	//fout.WriteString("grpc start 3")
	err = s.Serve(lis)
	library.Check(err, "grpc.server.serve error in web.GRPCServe")
	//fout.WriteString("grpc start 4")
}

type gRPCserver struct{}

func (grpc *gRPCserver) GetIP(ctx context.Context, r *proto.GetReq) (*proto.GetResp, error) {
	domain := r.Domain
	req := &model.GetReq{Domain: domain} //service
	resp := service.Srvc.GetIP(ctx, req) //service
	ret := &proto.GetResp{Domain: resp.Domain, IPs: resp.IPs}
	return ret, *new(error)
}

func (grpc *gRPCserver) Insert(ctx context.Context, r *proto.InsertReq) (*proto.InsertResp, error) {
	domain := r.Domain
	ip := r.IP
	req := &model.InsertReq{Domain: domain, IP: ip} //service
	resp := service.Srvc.Insert(ctx, req)           //service
	ret := &proto.InsertResp{Domain: resp.Domain, IP: resp.IP, Result: resp.Result}
	return ret, *new(error)
}

func (grpc *gRPCserver) Update(ctx context.Context, r *proto.UpdateReq) (*proto.UpdateResp, error) {
	dmsrc := r.Domainsrc
	ipsrc := r.IPsrc
	dmdst := r.Domaindst
	ipdst := r.IPdst
	req := &model.UpdateReq{Domainsrc: dmsrc, IPsrc: ipsrc, Domaindst: dmdst, IPdst: ipdst} //service
	resp := service.Srvc.Update(ctx, req)                                                   //service
	ret := &proto.UpdateResp{Affected: int64(resp.Affected), Result: resp.Result}
	return ret, *new(error)
}

func (grpc *gRPCserver) Delete(ctx context.Context, r *proto.DeleteReq) (*proto.DeleteResp, error) {
	domain := r.Domain
	ip := r.IP
	req := &model.DeleteReq{Domain: domain, IP: ip}
	resp := service.Srvc.Delete(ctx, req)
	ret := &proto.DeleteResp{Affected: int64(resp.Affected), Result: resp.Result}
	return ret, *new(error)
}
