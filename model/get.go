package model

type GetReq struct {
	Domain string `json:"Domain,omitempty"`
}

type GetResp struct {
	Domain string   `json:"Domain,omitempty"`
	IPs    []string `json:"IPs,omitempty"`
}
