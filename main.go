package main

import (
	"MiniDNS2/client"
	"MiniDNS2/model"
	"MiniDNS2/service"
	"MiniDNS2/web"
	"time"
)

func main() {
	service.InitService()
	go web.HTTPServe(model.Port1)
	go web.GRPCServe(model.Port2)
	//time.Sleep(3*time.Second)
	go client.GRPCClient(model.Port2)

	for {
		time.Sleep(time.Hour)
	}
}