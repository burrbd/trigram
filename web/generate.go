package web

import "net/http"

func GenerateHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Write([]byte("To think it more than commonly anxious"))
}
