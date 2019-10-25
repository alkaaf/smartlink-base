package model

type Resgen struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type AndroidResult struct {
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
}
type OtpResult struct {
	Code   int         `json:"code"`
	Reason string      `json:"reason"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
type AuthResult struct {
	Code   int         `json:"code"`
	Reason string      `json:"reason"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
