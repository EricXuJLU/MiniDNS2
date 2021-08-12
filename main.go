package main

import (
	"MiniDNS2/model"
	"MiniDNS2/service"
	"MiniDNS2/web"
)

func main() {
	service.InitService()
	web.HTTPServe(model.Port1)
}