package model

type InsertReq struct {
	Domain string
	IP string
}

type InsertResp struct {
	Domain string
	IP string
	Result string
}
