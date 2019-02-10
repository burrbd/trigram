package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheekybits/is"

	"github.com/burrbd/trigram/web"
)

type mockTrigramLearner struct {
	LearnFunc func(words []string)
}

func (learner mockTrigramLearner) Learn(words []string) {
	if learner.LearnFunc == nil {
		return
	}
	learner.LearnFunc(words)
}

func TestLearnHandler(t *testing.T) {
	is := is.New(t)
	h := web.LearnHandler(mockTrigramLearner{})
	srv := httptest.NewServer(h)
	defer srv.Close()
	w := httptest.NewRecorder()
	req:= httptest.NewRequest(http.MethodPost, srv.URL,
		strings.NewReader("To be or not to be, that is the question"))
	req.Header.Set("Content-Type", "text/plain")

	h.ServeHTTP(w, req)

	is.Equal("", w.Body.String())
	is.Equal(http.StatusAccepted, w.Result().StatusCode)
}

func TestLearnHandlerOnlyAcceptsPost(t *testing.T) {
	is := is.New(t)
	h := web.LearnHandler(mockTrigramLearner{})
	srv := httptest.NewServer(h)
	defer srv.Close()
	for _, method := range []string{"GET", "PUT", "PATCH", "DELETE", "HEAD", "CONNECT", "OPTIONS", "TRACE"} {
		w := httptest.NewRecorder()
		req:= httptest.NewRequest(method, srv.URL, nil)
		req.Header.Set("Content-Type", "text/plain")

		h.ServeHTTP(w, req)

		is.Equal(http.StatusMethodNotAllowed, w.Result().StatusCode)
	}
}

func TestLearnHandlerCallsTrigramLearner(t *testing.T) {
	is := is.New(t)

	learnerInvoked := false
	var learnerArg []string
	mockLearner := mockTrigramLearner{
		LearnFunc: func(words []string) {
			learnerInvoked, learnerArg = true, words
	}}

	req, err := http.NewRequest(http.MethodPost, "a_url",
		strings.NewReader("To be or not to be, that is the question"))
	is.NoErr(err)
	req.Header.Set("Content-Type", "text/plain")

	h := web.LearnHandler(mockLearner)
	h.ServeHTTP(httptest.NewRecorder(), req)

	is.True(learnerInvoked)
	is.Equal([]string{"to", "be" ,"or", "not", "to", "be", "that", "is", "the", "question"}, learnerArg)
}

func TestLearnHandlerCallsTrigramLearnerWithLineBreaks(t *testing.T) {
	is := is.New(t)

	learnerInvoked := false
	var learnerArg []string
	mockLearner := mockTrigramLearner{
		LearnFunc: func(words []string) {
			learnerInvoked, learnerArg = true, words
		}}

	req, err := http.NewRequest(http.MethodPost, "a_url",
		strings.NewReader(`
To be or
not to be,
that is the question`))

	is.NoErr(err)
	req.Header.Set("Content-Type", "text/plain")

	h := web.LearnHandler(mockLearner)
	h.ServeHTTP(httptest.NewRecorder(), req)

	is.True(learnerInvoked)
	is.Equal([]string{"to", "be" ,"or", "not", "to", "be", "that", "is", "the", "question"}, learnerArg)
}
