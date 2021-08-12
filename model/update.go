package model

type UpdateReq struct {
	Domainsrc string
	IPsrc string
	Domaindst string
	IPdst string
}

type UpdateResp struct {
	Affected int
	Result string
}
