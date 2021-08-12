package web

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"MiniDNS2/service"
	"context"
	"fmt"
	"net/http"
)


func HTTPServe(add string){
	http.HandleFunc("/get", GetIP)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	err := http.ListenAndServe(add, nil)
	library.Check(err, "HTTP.ListenAndServe error in web.HTTPServe")
}

func GetIP(w http.ResponseWriter, r *http.Request)  {
	domain :=r.URL.Query().Get("domain")
	req := &model.GetReq{Domain: domain}
	resp := service.Srvs.GetIP(context.Background(), req)
	if len(resp.IPs) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No sucn domain: %q\n", domain)
		return
	}else {
		fmt.Fprintf(w, "Domain:\n%q\nIP:\n", resp.Domain)
		for _, i := range resp.IPs {
			fmt.Fprintf(w, "%q\n", i)
		}
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	ip := r.URL.Query().Get("ip")
	if domain == "" || !library.IsIP(ip) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不合理的请求\n")
	}else {
		req := &model.InsertReq{
			Domain: domain,
			IP:     ip,
		}
		resp := service.Srvs.Insert(context.Background(), req)
		fmt.Fprintf(w, "%s\n", resp.Result)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	dmsrc := r.URL.Query().Get("dmsrc")
	ipsrc := r.URL.Query().Get("ipsrc")
	dmdst := r.URL.Query().Get("dmdst")
	ipdst := r.URL.Query().Get("ipdst")
	if dmsrc=="" || !library.IsIP(ipsrc) || dmdst=="" || !library.IsIP(ipdst) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不合理的请求")
	}else {
		req := &model.UpdateReq{
			Domainsrc: dmsrc,
			IPsrc:     ipsrc,
			Domaindst: dmdst,
			IPdst:     ipdst,
		}
		resp := service.Srvs.Update(context.Background(), req)
		fmt.Fprintf(w, resp.Result)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
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
		resp := service.Srvs.Delete(context.Background(), req)
		fmt.Fprintf(w, resp.Result)
	}
}