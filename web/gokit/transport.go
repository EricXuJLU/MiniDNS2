package gokit

import (
	"MiniDNS2/model"
	"context"
	"encoding/json"
	"net/http"
)

// func DecodeRequest(c context.Context, request *http.Request) (interface{}, error)
// func EncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error
func GetIPDecodeRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	domain := request.URL.Query().Get("domain")
	return model.GetReq{Domain: domain}, nil
}

func GetIPEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	//w.Header().Set()
	return json.NewEncoder(w).Encode(response)
}

func InsertDecodeRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	domain := request.FormValue("domain")
	ip := request.FormValue("ip")
	return model.InsertReq{Domain: domain, IP: ip}, nil
}

func InsertEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func UpdateDecodeRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	dmsrc := request.FormValue("dmsrc")
	ipsrc := request.FormValue("ipsrc")
	dmdst := request.FormValue("dmdst")
	ipdst := request.FormValue("ipdst")
	return model.UpdateReq{
		Domainsrc: dmsrc,
		IPsrc:     ipsrc,
		Domaindst: dmdst,
		IPdst:     ipdst,
	}, nil
}

func UpdateEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DeleteDecodeRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	domain := request.URL.Query().Get("domain")
	ip := request.URL.Query().Get("ip")
	return model.DeleteReq{Domain: domain, IP: ip}, nil
}

func DeleteEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
