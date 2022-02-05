package utils

import (
	"net/http"
)

func MessageHandler(w http.ResponseWriter, res []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(code))
	w.Write(res)
}
