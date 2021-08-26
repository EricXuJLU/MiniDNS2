package model

type UpdateReq struct {
	Domainsrc string `json:"Domainsrc,omitempty"`
	IPsrc     string `json:"IPsrc,omitempty"`
	Domaindst string `json:"Domaindst,omitempty"`
	IPdst     string `json:"IPdst,omitempty"`
}

type UpdateResp struct {
	Affected int    `json:"Affected,omitempty"`
	Result   string `json:"Result,omitempty"`
}
