package web

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"MiniDNS2/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HTTPServe(port string) {
	//被拦截的请求不写入日志
	http.HandleFunc("/getip", service.HttpChain(HTTPGetIP, service.HttpInterceptor("GET"), service.HttpLogger()))
	http.HandleFunc("/insert", service.HttpChain(HTTPInsert, service.HttpInterceptor("POST"), service.HttpLogger()))
	http.HandleFunc("/update", service.HttpChain(HTTPUpdate, service.HttpInterceptor("PUT"), service.HttpLogger()))
	http.HandleFunc("/delete", service.HttpChain(HTTPDelete, service.HttpInterceptor("DELETE"), service.HttpLogger()))
	err := http.ListenAndServe(port, nil)
	library.Check(err, "HTTP.ListenAndServe error in web.HTTPServe")
}

func HTTPGetIP(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	req := &model.GetReq{Domain: domain}
	resp := service.Srvc.GetIP(context.Background(), req)
	if len(resp.IPs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No sucn domain: %q\n", domain)
		return
	} else {
		fmt.Fprintf(w, "Domain:\n%q\nIP:\n", resp.Domain)
		for _, i := range resp.IPs {
			fmt.Fprintf(w, "%q\n", i)
		}
	}
}

func HTTPInsert(w http.ResponseWriter, r *http.Request) { //接受JSON输入
	data, err := ioutil.ReadAll(r.Body)
	library.Check(err, "ioutil.ReadAll error in web.httpServer.HTTPInsert")
	var value = make(map[string]interface{})
	err = json.Unmarshal(data, &value)
	library.Check(err, "json.Unmarshal error in web.httpServer.HTTPInsert")
	domain := value["domain"].(string)
	ip := value["ip"].(string)
	if domain == "" || !library.IsIP(ip) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不合理的请求\n%s\n%s", domain, ip)
	} else {
		req := &model.InsertReq{
			Domain: domain,
			IP:     ip,
		}
		resp := service.Srvc.Insert(context.Background(), req)
		fmt.Fprintf(w, "%s\n", resp.Result)
	}
}

func HTTPUpdate(w http.ResponseWriter, r *http.Request) { //接受JSON输入
	data, err := ioutil.ReadAll(r.Body)
	library.Check(err, "ioutil.ReadAll error in web.httpServer.HTTPInsert")
	var value = make(map[string]interface{})
	err = json.Unmarshal(data, &value)
	library.Check(err, "json.Unmarshal error in web.httpServer.HTTPInsert")
	dmsrc := value["dmsrc"].(string)
	ipsrc := value["ipsrc"].(string)
	dmdst := value["dmdst"].(string)
	ipdst := value["ipdst"].(string)
	if dmsrc == "" || !library.IsIP(ipsrc) || dmdst == "" || !library.IsIP(ipdst) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不合理的请求")
	} else {
		req := &model.UpdateReq{
			Domainsrc: dmsrc,
			IPsrc:     ipsrc,
			Domaindst: dmdst,
			IPdst:     ipdst,
		}
		resp := service.Srvc.Update(context.Background(), req)
		fmt.Fprintf(w, resp.Result)
	}
}

func HTTPDelete(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	ip := r.URL.Query().Get("ip")
	if domain == "" || !library.IsIP(ip) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不合理的请求")
	} else {
		req := &model.DeleteReq{
			Domain: domain,
			IP:     ip,
		}
		resp := service.Srvc.Delete(context.Background(), req)
		fmt.Fprintf(w, resp.Result)
	}
}
