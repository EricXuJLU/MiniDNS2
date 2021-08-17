package web

import (
	"MiniDNS2/service"
	"MiniDNS2/web/gokit"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func GokitServe(port string) {
	gokitServeMux := http.NewServeMux()
	gk := gokit.NewEndpoints(service.Srvc)
	var gksvr *kithttp.Server
	gksvr = kithttp.NewServer(gk.GetIPEndpoint, gokit.GetIPDecodeRequest, gokit.GetIPEncodeResponse)
	gokitServeMux.Handle("/getip", gksvr)
	gksvr = kithttp.NewServer(gk.InsertEndpoint, gokit.InsertDecodeRequest, gokit.InsertEncodeResponse)
	gokitServeMux.Handle("/insert", gksvr)
	gksvr = kithttp.NewServer(gk.UpdateEndpoint, gokit.UpdateDecodeRequest, gokit.UpdateEncodeResponse)
	gokitServeMux.Handle("/update", gksvr)
	gksvr = kithttp.NewServer(gk.DeleteEndpoint, gokit.DeleteDecodeRequest, gokit.DeleteEncodeResponse)
	gokitServeMux.Handle("/delete", gksvr)
	http.ListenAndServe(port, gokitServeMux)
}
