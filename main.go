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
	go web.HTTPServe(model.Port1) //10086
	//gRPC
	go web.GRPCServe(model.Port2) //10010
	go client.GRPCClient(model.Port2)
	//Gin
	go web.GinServe(model.Port3) //3985
	//gokit
	go web.GokitServe(model.Port4) //2021
	//grpc-gateway-http
	go web.GatewayServe(model.Port5, model.Local+model.Port2) //2077
	//主进程睡眠
	for {
		time.Sleep(time.Hour)
	}
}
