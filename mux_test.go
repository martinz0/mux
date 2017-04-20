package mux

import (
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	router := New()
	router.Get("/v1/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(r.Context().Value("_param").(*params).Get([]byte("id")))
	}))
	http.ListenAndServe(":1234", router)
}

func BenchmarkMux(b *testing.B) {
	b.ReportAllocs()
	router := New()
	router.Get("/v1/:id", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	req, _ := http.NewRequest("GET", "/v1/123", nil)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, req)
	}
}
