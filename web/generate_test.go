package web_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram/web"
)

func TestGenerateHandler(t *testing.T) {
	is := is.New(t)

	h := http.HandlerFunc(web.GenerateHandler)

	srv := httptest.NewServer(h)
	defer srv.Close()

	req := httptest.NewRequest("GET", srv.URL, nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	expBody := "To think it more than commonly anxious"
	actBody := w.Body.String()
	is.Equal(expBody, actBody)
	is.Equal(http.StatusOK, w.Result().StatusCode)
}
