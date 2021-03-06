package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{
			in:  "dennis@golang.org",
			out: "$$$",
		},
		{
			in:  "dennis",
			out: "Hello, World!",
		},
	}
	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("expected response body to be %s, got %s", c.out, rec.Body.String())
		}
	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/dennis@golang.org",
			nil,
		)
		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("expected status code 200, got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "$$$") {
			b.Errorf("expected response body to be %s, got %s", "$$$", rec.Body.String())
		}
	}
}
