package web_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram/web"
)

type mockNaturalLanguageGenerator struct {
	GenerateFunc func() string
}

func (g mockNaturalLanguageGenerator) Generate() string {
	if g.GenerateFunc == nil {
		return "To think it more than commonly anxious"
	}
	return g.GenerateFunc()
}

func TestGenerateHandler(t *testing.T) {
	is := is.New(t)

	h := web.GenerateHandler(mockNaturalLanguageGenerator{})

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

func TestGenerateHandlerOnlyAcceptsGet(t *testing.T) {
	is := is.New(t)
	h := web.GenerateHandler(mockNaturalLanguageGenerator{})
	srv := httptest.NewServer(h)
	defer srv.Close()
	for _, method := range []string{"POST", "PUT", "PATCH", "DELETE", "HEAD", "CONNECT", "OPTIONS", "TRACE"} {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(method, srv.URL, nil)

		h.ServeHTTP(w, req)

		is.Equal(http.StatusMethodNotAllowed, w.Result().StatusCode)
	}
}

func TestGenerateHandlerCallsLanguageGenerator(t *testing.T) {
	is := is.New(t)

	generatorInvoked := false
	expBody := "Their visit afforded was produced by the lady with whom she almost looked up to the stables."
	mockGenerator := mockNaturalLanguageGenerator{
		GenerateFunc: func() string {
			generatorInvoked = true
			return expBody
		}}

	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "a_url", nil)
	is.NoErr(err)
	req.Header.Set("Content-Type", "text/plain")

	h := web.GenerateHandler(mockGenerator)
	h.ServeHTTP(w, req)

	is.True(generatorInvoked)
	is.Equal(expBody, w.Body.String())
}
