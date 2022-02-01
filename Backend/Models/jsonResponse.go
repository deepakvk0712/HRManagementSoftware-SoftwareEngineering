package models

type JsonResponse struct {
	Error string `json:"error"`
	Msg   string `json:"msg"`
	Data  string `json:"data"`
}
