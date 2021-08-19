package web

import (
	"MiniDNS2/model"
	"MiniDNS2/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServe(port string) {
	router := gin.Default()
	router.GET("/", Index_api) //欢迎页
	router.GET("/getip", GetIP_api)
	router.POST("/insert", Insert_api)
	router.PUT("/update", Update_api)
	router.DELETE("/delete", Delete_api)
	router.Run(port)
}

func Index_api(c *gin.Context) {
	c.String(http.StatusOK, model.GinIndex)
}

func GetIP_api(c *gin.Context) {
	domain := c.Request.URL.Query().Get("domain")
	req := &model.GetReq{Domain: domain}
	resp := service.Srvc.GetIP(context.Background(), req)
	c.JSON(http.StatusOK, resp)
}

func Insert_api(c *gin.Context) {
	domain := c.Request.FormValue("domain")
	ip := c.Request.FormValue("ip")
	req := &model.InsertReq{Domain: domain, IP: ip}
	resp := service.Srvc.Insert(context.Background(), req)
	c.JSON(http.StatusOK, resp)
}

func Update_api(c *gin.Context) {
	dmsrc := c.Request.FormValue("dmsrc")
	ipsrc := c.Request.FormValue("ipsrc")
	dmdst := c.Request.FormValue("dmdst")
	ipdst := c.Request.FormValue("ipdst")
	req := &model.UpdateReq{
		Domainsrc: dmsrc,
		IPsrc:     ipsrc,
		Domaindst: dmdst,
		IPdst:     ipdst,
	}
	resp := service.Srvc.Update(context.Background(), req)
	c.JSON(http.StatusOK, resp)
}

func Delete_api(c *gin.Context) {
	domain := c.Request.URL.Query().Get("domain")
	ip := c.Request.URL.Query().Get("ip")
	req := &model.DeleteReq{Domain: domain, IP: ip}
	resp := service.Srvc.Delete(context.Background(), req)
	c.JSON(http.StatusOK, resp)
}
