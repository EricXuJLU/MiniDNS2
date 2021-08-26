package transport

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// func DecodeRequest(c context.Context, request *http.Request) (interface{}, error)
// func EncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error

func GetIPDecodeRequest(_ context.Context, request *http.Request) (interface{}, error) {
	domain := request.URL.Query().Get("domain")
	return model.GetReq{Domain: domain}, nil
}

func GetIPEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	//w.Header().Set()
	return json.NewEncoder(w).Encode(response)
}

func InsertDecodeRequest(_ context.Context, request *http.Request) (interface{}, error) {
	data, err := ioutil.ReadAll(request.Body)
	library.Check(err, "ioutil.ReadAll error in web.gokit.transport.InsertDecodeRequest")
	var value = make(map[string]interface{})
	err = json.Unmarshal(data, &value)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.InsertDecodeRequest")
	domain := value["domain"].(string)
	ip := value["ip"].(string)
	return model.InsertReq{Domain: domain, IP: ip}, nil
}

func InsertEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func UpdateDecodeRequest(_ context.Context, request *http.Request) (interface{}, error) {
	data, err := ioutil.ReadAll(request.Body)
	library.Check(err, "ioutil.ReadAll error in web.gokit.transport.UpdateDecodeRequest")
	var value = make(map[string]interface{})
	err = json.Unmarshal(data, &value)
	library.Check(err, "json.Unmarshal error in web.gokit.transport.UpdateDecodeRequest")
	dmsrc := value["dmsrc"].(string)
	ipsrc := value["ipsrc"].(string)
	dmdst := value["dmdst"].(string)
	ipdst := value["ipdst"].(string)
	return model.UpdateReq{
		Domainsrc: dmsrc,
		IPsrc:     ipsrc,
		Domaindst: dmdst,
		IPdst:     ipdst,
	}, nil
}

func UpdateEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func DeleteDecodeRequest(_ context.Context, request *http.Request) (interface{}, error) {
	domain := request.URL.Query().Get("domain")
	ip := request.URL.Query().Get("ip")
	return model.DeleteReq{Domain: domain, IP: ip}, nil
}

func DeleteEncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
