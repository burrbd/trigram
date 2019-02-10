package web

import "net/http"

func GenerateHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("To think it more than commonly anxious"))
}