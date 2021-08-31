package transport

import (
	"MiniDNS2/model"
	"MiniDNS2/proto"
	"context"
	"log"
	"reflect"
	"testing"
)

func TestDecodeGRPCUpdateRequest(t *testing.T) {
	get, err := DecodeGRPCUpdateRequest(context.Background(), &proto.UpdateReq{
		Domainsrc: "google.com",
		IPsrc:     "1.1.1.1",
		Domaindst: "test.com",
		IPdst:     "8.8.8.8",
	})
	if err != nil {
		log.Fatal(err)
	}
	want := model.UpdateReq{
		Domainsrc: "google.com",
		IPsrc:     "1.1.1.1",
		Domaindst: "test.com",
		IPdst:     "8.8.8.8",
	}
	if !reflect.DeepEqual(get, want) {
		t.Fatal("Get", get, "while want", want)
	}
}
