package web

import (
	"MiniDNS2/library"
	"MiniDNS2/model"
	"MiniDNS2/service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func GinServe(port string) {
	router := gin.New()
	router.Use(service.GinLogger())
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
	data, err := ioutil.ReadAll(c.Request.Body)
	library.Check(err, "ioutil.ReadAll error in web.ginServer.Insert_api")
	var value = make(map[string]interface{})
	err = json.Unmarshal(data, &value)
	library.Check(err, "json.Unmarshal error in web.ginServer.Insert_api")
	domain := value["domain"].(string)
	ip := value["ip"].(string)
	req := &model.InsertReq{Domain: domain, IP: ip}
	resp := service.Srvc.Insert(context.Background(), req)
	c.JSON(http.StatusOK, resp)
}

func Update_api(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	library.Check(err, "ioutil.ReadAll error in web.ginServer.Update_api")
	var value = make(map[string]interface{})
	err = json.Unmarshal(data, &value)
	library.Check(err, "json.Unmarshal error in web.ginServer.Update_api")
	dmsrc := value["dmsrc"].(string)
	ipsrc := value["ipsrc"].(string)
	dmdst := value["dmdst"].(string)
	ipdst := value["ipdst"].(string)
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
