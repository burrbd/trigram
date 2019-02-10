package web

import "net/http"

func LearnHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)
}