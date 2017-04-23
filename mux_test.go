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
	initRoute(router)
	fmt.Println(router.entry)
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
}

func BenchmarkMux_1(b *testing.B) {
	b.ReportAllocs()
	router := New()
	initRoute(router)
	req, _ := http.NewRequest("GET", "/lor/v1/users", nil)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, req)
	}
}

func BenchmarkMux_2(b *testing.B) {
	b.ReportAllocs()
	router := New()
	initRoute(router)
	req, _ := http.NewRequest("GET", "/lor/v1/users/1/classes/2/teachers/3", nil)
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

var testRouterCase = struct {
	methods []string
	paths   []string
}{
	[]string{"GET", "POST", "PUT", "DELETE"},
	[]string{
		"/",
		"/lor",
		"/lor/v1",
		"/lor/v1/users",
		"/lor/v1/users/:uid",
		"/lor/v1/users/:uid/name",
		"/lor/v1/users/:uid/habbit",
		"/lor/v1/users/:uid/classes",
		"/lor/v1/users/:uid/classes/:classid",
		"/lor/v1/users/:uid/classes/:classid/teachers",
		"/lor/v1/users/:uid/classes/:classid/teachers/:teacherid",
		"/lor/v1/users/:uid/classes/:classid/teachers/:teacherid/classes",
		"/lor/v1/users/:uid/classes/:classid/teachers/:teacherid/classes/:classid",
	},
}

func initRoute(r *Mux) {
	for _, method := range testRouterCase.methods {
		for _, path := range testRouterCase.paths {
			r.Handle(method, path, func(w http.ResponseWriter, r *http.Request, ps Params) {})
		}
	}
}
