package model

type Resgen struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type AndroidResult struct {
	Msg    string `json:"msg"`
	Reason string `json:"reason"`
}
