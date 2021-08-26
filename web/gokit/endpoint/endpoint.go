package endpoint

import (
	"MiniDNS2/model"
	"MiniDNS2/service"
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetIPEndpoint  endpoint.Endpoint
	InsertEndpoint endpoint.Endpoint
	UpdateEndpoint endpoint.Endpoint
	DeleteEndpoint endpoint.Endpoint
}

func NewEndpoints(srv *service.Service) *Endpoints {
	return &Endpoints{
		GetIPEndpoint:  service.GoKitLogger(MakeGetIPEndpoint(srv)),
		InsertEndpoint: service.GoKitLogger(MakeInsertEndpoint(srv)),
		UpdateEndpoint: service.GoKitLogger(MakeUpdateEndpoint(srv)),
		DeleteEndpoint: service.GoKitLogger(MakeDeleteEndpoint(srv)),
	}
}

func MakeGetIPEndpoint(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.GetReq)
		return srv.GetIP(ctx, &req), nil
	}
}

func MakeInsertEndpoint(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.InsertReq)
		return srv.Insert(ctx, &req), nil
	}
}

func MakeUpdateEndpoint(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.UpdateReq)
		return srv.Update(ctx, &req), nil
	}
}

func MakeDeleteEndpoint(srv *service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(model.DeleteReq)
		return srv.Delete(ctx, &req), nil
	}
}
