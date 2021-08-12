package model

type DeleteReq struct {
	Domain string
	IP string
}

type DeleteResp struct {
	Affected int
	Result string
}
