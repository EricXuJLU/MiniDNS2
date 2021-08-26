package transport

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"MiniDNS2/proto"
	"MiniDNS2/web/gokit/endpoint"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/transport/grpc"
)

//封装在Gokit里的grpc，不直接与Service层交互
/*
type Handler interface {
	ServeGPRC(ctx context.Context, req interface{}) (context.Context, interface{}, error)
}

type Server struct {
	e	kitendpoint.Endpoint
	dec grpc.DecodeRequestFunc
	enc grpc.EncodeResponseFunc
}

func NewServer(
	e kitendpoint.Endpoint,
	dec grpc.DecodeRequestFunc,
	enc grpc.EncodeResponseFunc,
	options ...grpc.ServerOption,
	) *Server {
	s := &Server{
		e:   e,
		dec: dec,
		enc: enc,
	}
	for _, option := range options {
		option(s)
	}
}
*/

type KitGRPCServer struct {
	getIP  grpc.Handler
	insert grpc.Handler
	update grpc.Handler
	delete grpc.Handler
}

func NewKitGRPCServer(endpoints *endpoint.Endpoints) *KitGRPCServer {
	getIPHandler := grpc.NewServer(endpoints.GetIPEndpoint, DecodeGRPCGetIPRequest, EncodeGRPCGetIPResponse)
	insertHandler := grpc.NewServer(endpoints.InsertEndpoint, DecodeGRPCInsertRequest, EncodeGRPCInsertResponse)
	updateHandler := grpc.NewServer(endpoints.UpdateEndpoint, DecodeGRPCUpdateRequest, EncodeGRPCUpdateResponse)
	deleteHandler := grpc.NewServer(endpoints.DeleteEndpoint, DecodeGRPCDeleteRequest, EncodeGRPCDeleteResponse)
	s := &KitGRPCServer{
		getIP:  getIPHandler,
		insert: insertHandler,
		update: updateHandler,
		delete: deleteHandler,
	}
	return s
}

func (s *KitGRPCServer) GetIP(ctx context.Context, req *proto.GetReq) (*proto.GetResp, error) {
	_, resp, err := s.getIP.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.GetResp), nil
}

func (s *KitGRPCServer) Insert(ctx context.Context, req *proto.InsertReq) (*proto.InsertResp, error) {
	_, resp, err := s.insert.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.InsertResp), nil
}

func (s *KitGRPCServer) Update(ctx context.Context, req *proto.UpdateReq) (*proto.UpdateResp, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.UpdateResp), nil
}

func (s *KitGRPCServer) Delete(ctx context.Context, req *proto.DeleteReq) (*proto.DeleteResp, error) {
	_, resp, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.DeleteResp), nil
}

//参数req是grpc的，返回值是model层的，都是指针
func DecodeGRPCGetIPRequest(_ context.Context, req interface{}) (interface{}, error) {
	Req := req.(*proto.GetReq)
	bt, err := json.Marshal(Req)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.DecodeGRPCGetIPRequest")
	var ret model.GetReq
	err = json.Unmarshal(bt, &ret)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.DecodeGRPCGetIPRequest")
	return ret, err
}

//参数resp是model层的，返回值是grpc的，都是指针
func EncodeGRPCGetIPResponse(_ context.Context, resp interface{}) (interface{}, error) {
	var data *proto.GetResp
	Resp := resp.(*model.GetResp)
	bt, err := json.Marshal(Resp)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.EncodeGRPCGetIPResponse")
	err = json.Unmarshal(bt, &data)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.EncodeGRPCGetIPResponse")
	return data, err
}

//参数req是grpc的，返回值是model层的，都是指针
func DecodeGRPCInsertRequest(_ context.Context, req interface{}) (interface{}, error) {
	Req := req.(*proto.InsertReq)
	bt, err := json.Marshal(Req)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.DecodeGRPCInsertRequest")
	var ret model.InsertReq
	err = json.Unmarshal(bt, &ret)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.DecodeGRPCInsertRequest")
	return ret, err
}

//参数resp是model层的，返回值是grpc的，都是指针
func EncodeGRPCInsertResponse(_ context.Context, resp interface{}) (interface{}, error) {
	var data *proto.InsertResp
	Resp := resp.(*model.InsertResp)
	bt, err := json.Marshal(Resp)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.EncodeGRPCInsertResponse")
	err = json.Unmarshal(bt, &data)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.EncodeGRPCInsertResponse")
	return data, err
}

//参数req是grpc的，返回值是model层的，都是指针
func DecodeGRPCUpdateRequest(_ context.Context, req interface{}) (interface{}, error) {
	Req := req.(*proto.UpdateReq)
	bt, err := json.Marshal(Req)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.DecodeGRPCUpdateRequest")
	var ret model.UpdateReq
	err = json.Unmarshal(bt, &ret)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.DecodeGRPCUpdateRequest")
	return ret, err
}

//参数resp是model层的，返回值是grpc的，都是指针
func EncodeGRPCUpdateResponse(_ context.Context, resp interface{}) (interface{}, error) {
	var data *proto.UpdateResp
	Resp := resp.(*model.UpdateResp)
	bt, err := json.Marshal(Resp)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.EncodeGRPCUpdateResponse")
	err = json.Unmarshal(bt, &data)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.EncodeGRPCUpdateResponse")
	return data, err
}

//参数req是grpc的，返回值是model层的，都是指针
func DecodeGRPCDeleteRequest(_ context.Context, req interface{}) (interface{}, error) {
	Req := req.(*proto.DeleteReq)
	bt, err := json.Marshal(Req)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.DecodeGRPCDeleteRequest")
	var ret model.DeleteReq
	err = json.Unmarshal(bt, &ret)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.DecodeGRPCDeleteRequest")
	return ret, err
}

//参数resp是model层的，返回值是grpc的，都是指针
func EncodeGRPCDeleteResponse(_ context.Context, resp interface{}) (interface{}, error) {
	var data *proto.DeleteResp
	Resp := resp.(*model.DeleteResp)
	bt, err := json.Marshal(Resp)
	library.Check(err, "json.Marshal error in web.gokit.transport.grpc.EncodeGRPCDeleteResponse")
	err = json.Unmarshal(bt, &data)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.grpc.EncodeGRPCDeleteResponse")
	return data, err
}
