package mux

import (
	"context"
	"net/http"
)

type Mux struct {
	entry *muxEntry
}

func New() *Mux {
	return &Mux{
		entry: NewMuxEntry(),
	}
}

func (m *Mux) Get(path string, handler http.Handler) {
	m.handle("GET", path, handler)
}

func (m *Mux) Post(path string, handler http.Handler) {
	m.handle("POST", path, handler)
}

func (m *Mux) Put(path string, handler http.Handler) {
	m.handle("PUT", path, handler)
}

func (m *Mux) Delete(path string, handler http.Handler) {
	m.handle("DELETE", path, handler)
}

func (m *Mux) handle(method, path string, handler http.Handler) {
	m.entry.Add([]byte(method), []byte(path), handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var param params
	handler := m.entry.Lookup([]byte(r.Method), []byte(r.URL.Path), &param)
	if len(param.params) > 0 {
		ctx := context.WithValue(r.Context(), "_param", &param)
		r = r.WithContext(ctx)
	}
	handler.ServeHTTP(w, r)
}
