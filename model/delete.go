package model

type DeleteReq struct {
	Domain string `json:"Domain,omitempty"`
	IP     string `json:"IP,omitempty"`
}

type DeleteResp struct {
	Affected int    `json:"Affected,omitempty"`
	Result   string `json:"Result,omitempty"`
}
