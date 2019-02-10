package web_test

import (
	"github.com/burrbd/trigram/web"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheekybits/is"
)

func TestLearnPostHandler(t *testing.T) {
	is := is.New(t)
	h := http.HandlerFunc(web.LearnHandler)
	srv := httptest.NewServer(h)
	defer srv.Close()
	w := httptest.NewRecorder()
	req:= httptest.NewRequest("POST", srv.URL,
		strings.NewReader("To be or not to be, that is the question"))
	req.Header.Set("Content-Type", "text/plain")


	h.ServeHTTP(w, req)

	n, _ := w.Result().Body.Read(make([]byte, 1))
	is.Equal(0, n)
	is.Equal(http.StatusAccepted, w.Result().StatusCode)
}
