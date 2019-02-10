package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram/web"
)

func TestLearnPostHandler(t *testing.T) {
	is := is.New(t)
	h := http.HandlerFunc(web.LearnHandler)
	srv := httptest.NewServer(h)
	defer srv.Close()
	w := httptest.NewRecorder()
	req:= httptest.NewRequest(http.MethodPost, srv.URL,
		strings.NewReader("To be or not to be, that is the question"))
	req.Header.Set("Content-Type", "text/plain")

	h.ServeHTTP(w, req)

	n, _ := w.Result().Body.Read(make([]byte, 1))
	is.Equal(0, n)
	is.Equal(http.StatusAccepted, w.Result().StatusCode)
}

func TestLearnHandlerOnlyAcceptsPost(t *testing.T) {
	is := is.New(t)
	h := http.HandlerFunc(web.LearnHandler)
	srv := httptest.NewServer(h)
	defer srv.Close()
	for _, method := range []string{"GET", "PUT", "PATCH", "DELETE", "HEAD", "CONNECT", "OPTIONS", "TRACE"} {
		w := httptest.NewRecorder()
		req:= httptest.NewRequest(method, srv.URL, nil)
		req.Header.Set("Content-Type", "text/plain")

		h.ServeHTTP(w, req)

		is.Equal(http.StatusBadRequest, w.Result().StatusCode)
	}
}
