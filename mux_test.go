package mux

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMux(t *testing.T) {
	router := New()
	router.Get("/v1/*", func(w http.ResponseWriter, r *http.Request, ps params) {
		w.Write([]byte("ok"))
	})
	req, _ := http.NewRequest("GET", "/v1/123", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
}

func BenchmarkMux(b *testing.B) {
	b.ReportAllocs()
	router := New()
	router.Get("/v1/:id", func(w http.ResponseWriter, r *http.Request, ps params) {
	})
	req, _ := http.NewRequest("GET", "/v1/123", nil)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, req)
	}
}
