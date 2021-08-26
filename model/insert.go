package model

type InsertReq struct {
	Domain string `json:"Domain,omitempty"`
	IP     string `json:"IP,omitempty"`
}

type InsertResp struct {
	Domain string `json:"Domain,omitempty"`
	IP     string `json:"IP,omitempty"`
	Result string `json:"Result,omitempty"`
}
