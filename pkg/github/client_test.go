package github

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shurcooL/githubv4"
)

func mustRead(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func mustWrite(w io.Writer, s string) {
	_, err := io.WriteString(w, s)
	if err != nil {
		panic(err)
	}
}

// https://github.com/shurcooL/githubv4/blob/master/githubv4_test.go
// localRoundTripper is an http.RoundTripper that executes HTTP transactions
// by using handler directly, instead of going over an HTTP connection.
type localRoundTripper struct {
	handler http.Handler
}

func (l localRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	w := httptest.NewRecorder()
	l.handler.ServeHTTP(w, req)
	return w.Result(), nil
}

func mockClient(input, output string, t *testing.T) *client {
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", func(w http.ResponseWriter, req *http.Request) {
		if got, want := req.Method, http.MethodPost; got != want {
			t.Errorf("got request method: %v, want: %v", got, want)
		}
		body := mustRead(req.Body)
		if got, want := body, input+"\n"; got != want {
			// t.Errorf("got body: %v, want %v", got, want)
		}
		w.Header().Set("Content-Type", "application/json")
		mustWrite(w, output)
	})
	c := githubv4.NewClient(&http.Client{Transport: localRoundTripper{handler: mux}})
	return &client{v4Cli: c}
}

func TestGetMeta(t *testing.T) {
	v4 := mockClient("", "", t)
	v4.GetMeta("", "")
}
