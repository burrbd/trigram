package web

import (
	"net/http"

	"github.com/burrbd/trigram"
)

func GenerateHandler(generator trigram.LanguageGenerator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		w.Write([]byte(generator.Generate()))
	})
}
