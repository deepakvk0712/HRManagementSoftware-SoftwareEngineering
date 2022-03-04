package models

type JsonResponseObject struct {
	Error string      `json:"error"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}
