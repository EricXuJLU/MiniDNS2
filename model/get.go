package model

type GetReq struct {
	Domain string
}

type GetResp struct {
	Domain string
	IPs []string
}
