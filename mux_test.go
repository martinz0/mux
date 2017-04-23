package mux

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestMux(t *testing.T) {
	router := New()
	initRoute(router)
	req, _ := http.NewRequest("GET", "/lor/v1/users/1000", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
	req, _ = http.NewRequest("GET", "/lor/v1/users/1000/classes", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
	req, _ = http.NewRequest("GET", "/lor/v1/users/1000/classes/1000", nil)
	router.ServeHTTP(httptest.NewRecorder(), req)
	req, _ = http.NewRequest("GET", "/lor/v1/users/1000/classes/1000/teachers", nil)
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

func BenchmarkMux_3(b *testing.B) {
	b.ReportAllocs()
	router := New()
	n := initRoute(router)
	for i := 0; i < b.N; i++ {
		router.ServeHTTP(nil, reqs[i%n])
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
		"/lor/v2",
		"/lor/v2/users",
		"/lor/v2/users/:uid",
		"/lor/v2/users/:uid/name",
		"/lor/v2/users/:uid/habbit",
		"/lor/v2/users/:uid/classes",
		"/lor/v2/users/:uid/classes/:classid",
		"/lor/v2/users/:uid/classes/:classid/teachers",
		"/lor/v2/users/:uid/classes/:classid/teachers/:teacherid",
		"/lor/v2/users/:uid/classes/:classid/teachers/:teacherid/classes",
		"/lor/v2/users/:uid/classes/:classid/teachers/:teacherid/classes/:classid",
		"/lor/v3",
		"/lor/v3/users",
		"/lor/v3/users/:uid",
		"/lor/v3/users/:uid/name",
		"/lor/v3/users/:uid/habbit",
		"/lor/v3/users/:uid/classes",
		"/lor/v3/users/:uid/classes/:classid",
		"/lor/v3/users/:uid/classes/:classid/teachers",
		"/lor/v3/users/:uid/classes/:classid/teachers/:teacherid",
		"/lor/v3/users/:uid/classes/:classid/teachers/:teacherid/classes",
		"/lor/v3/users/:uid/classes/:classid/teachers/:teacherid/classes/:classid",
		"/lor/v4",
		"/lor/v4/users",
		"/lor/v4/users/:uid",
		"/lor/v4/users/:uid/name",
		"/lor/v4/users/:uid/habbit",
		"/lor/v4/users/:uid/classes",
		"/lor/v4/users/:uid/classes/:classid",
		"/lor/v4/users/:uid/classes/:classid/teachers",
		"/lor/v4/users/:uid/classes/:classid/teachers/:teacherid",
		"/lor/v4/users/:uid/classes/:classid/teachers/:teacherid/classes",
		"/lor/v4/users/:uid/classes/:classid/teachers/:teacherid/classes/:classid",
		"/lor/v5",
		"/lor/v5/users",
		"/lor/v5/users/:uid",
		"/lor/v5/users/:uid/name",
		"/lor/v5/users/:uid/habbit",
		"/lor/v5/users/:uid/classes",
		"/lor/v5/users/:uid/classes/:classid",
		"/lor/v5/users/:uid/classes/:classid/teachers",
		"/lor/v5/users/:uid/classes/:classid/teachers/:teacherid",
		"/lor/v5/users/:uid/classes/:classid/teachers/:teacherid/classes",
		"/lor/v5/users/:uid/classes/:classid/teachers/:teacherid/classes/:classid",
	},
}

func init() {
	for _, method := range testRouterCase.methods {
		for _, path := range testRouterCase.paths {
			req, _ := http.NewRequest(method, path, nil)
			reqs = append(reqs, req)
		}
	}
}

var reqs []*http.Request

func initRoute(r *Mux) int {
	for _, method := range testRouterCase.methods {
		for _, path := range testRouterCase.paths {
			r.Handle(method, path, func(w http.ResponseWriter, r *http.Request, ps Params) {})
		}
	}
	return len(testRouterCase.methods) * len(testRouterCase.paths)
}
