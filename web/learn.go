package web

import (
	"bufio"
	"net/http"
	"strings"

	"github.com/burrbd/trigram"
)

func LearnHandler(learner trigram.Learner) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		words := make([]string, 0)
		scanner := bufio.NewScanner(req.Body)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words = append(words, strings.ToLower(
				strings.Trim(scanner.Text(), `'",.?!()[]{}`)))
		}
		learner.Learn(words)
		w.WriteHeader(http.StatusAccepted)
	})
}
