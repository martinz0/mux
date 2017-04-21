package mux

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMux(t *testing.T) {
	router := New()
	router.Handle("GET", "/", func(w http.ResponseWriter, r *http.Request, ps Params) {
	})
	fmt.Println(router.entry)
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
}

func BenchmarkMux(b *testing.B) {
	b.ReportAllocs()
	router := New()
	router.Handle("GET", "/v1/v1", func(w http.ResponseWriter, r *http.Request, ps Params) {
	})
	req, _ := http.NewRequest("GET", "/v1/v1", nil)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, req)
	}
}

func BenchmarkAddHandler(b *testing.B) {
	b.ReportAllocs()
	router := New()
	for i := 0; i < b.N; i++ {
		router.Handle("GET", strconv.Itoa(i), nil)
	}
}
