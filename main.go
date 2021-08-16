package main

import (
	"MiniDNS2/client"
	"MiniDNS2/model"
	"MiniDNS2/service"
	"MiniDNS2/web"
	"time"
)

func main() {
	//服务层初始化
	service.InitService()
	//http
	go web.HTTPServe(model.Port1)
	//gRPC
	go web.GRPCServe(model.Port2)
	go client.GRPCClient(model.Port2)
	//Gin
	go web.GinServe(model.Port3)
	//主进程睡眠
	for {
		time.Sleep(time.Hour)
	}
}
