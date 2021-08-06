package service

import (
	"MiniDNS2/dao"
	"MiniDNS2/library"
	"fmt"
	"net/http"
)

func HTTPServe(add string){
	http.HandleFunc("/get", GetIP)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	err := http.ListenAndServe(add, nil)
	library.Check(err, 2001)
}

func GetIP(w http.ResponseWriter, r *http.Request)  {
	domain :=r.URL.Query().Get("domain")
	ips := dao.GetIP(domain)
	if len(ips) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No sucn domain: %q\n", domain)
		return
	}else {
		fmt.Fprintf(w, "Domain:\n%q\nIP:\n", domain)
		for _, i := range ips {
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
		fmt.Fprintf(w, "%s\n", dao.Insert(domain, ip))
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
		aff := dao.Update(dmsrc, ipsrc, dmdst, ipdst)
		fmt.Fprintf(w, "%d个条目已被更新", aff)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	domain := r.URL.Query().Get("domain")
	ip := r.URL.Query().Get("ip")
	if domain == "" || !library.IsIP(ip) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不合理的请求")
	} else {
		aff := dao.Delete(domain, ip)
		fmt.Fprintf(w, "%d个条目已被删除", aff)
	}
}
