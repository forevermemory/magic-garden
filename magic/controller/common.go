package controller

// Response code 0 成功 code -1 失败
type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Errcode int    `json:"code"`
	Message string `json:"msg"`
}

type OKResponse struct {
	Errcode int         `json:"code"`
	Data    interface{} `json:"data"`
}
