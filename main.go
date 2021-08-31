package main

import (
	"MiniDNS2/model"
	"MiniDNS2/service"
	"MiniDNS2/web"
	"MiniDNS2/web/grpc"
	"MiniDNS2/web/grpc/client"
	"log"
	"time"
)

func init() {
	//服务层初始化
	service.InitService()
	//日志格式
	log.SetPrefix("Log: ")
	log.SetFlags(log.Lmicroseconds | log.Lshortfile)
}

func main() {
	//http
	go web.HTTPServe(model.Port1) //10086

	//gRPC
	go grpc.GRPCServe(model.Port2) //10010
	go client.GRPCClient(model.Port2)

	//Gin
	go web.GinServe(model.Port3) //3985

	//GoKit
	go web.GoKitServe(model.Port4) //2021

	//grpc-gateway-http
	go web.GatewayServe(model.Port5, model.Local+model.Port2) //2077

	//GoKit-gRPC-gateway
	go web.GoKitGRPCServe(model.Port6)                        //1573
	go web.GatewayServe(model.Port7, model.Local+model.Port6) //6114

	//主进程睡眠
	for {
		time.Sleep(time.Hour)
	}
}
