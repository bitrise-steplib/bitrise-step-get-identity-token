package api

type GetIdentityTokenParameter struct {
	Subject  string `json:"sub"`
	Audience string `json:"aud"`
}

type errorReponse struct {
	Message string `json:"error_msg"`
}
