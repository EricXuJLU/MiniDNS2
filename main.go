package main

import (
	"MiniDNS2/model"
	"MiniDNS2/service"
)

func main() {
	service.HTTPServe(model.Port1)
}