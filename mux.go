package mux

import (
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

type Handler func(w http.ResponseWriter, r *http.Request, ps params)

func (m *Mux) Get(path string, handler Handler) {
	m.handle("GET", path, handler)
}

func (m *Mux) Post(path string, handler Handler) {
	m.handle("POST", path, handler)
}

func (m *Mux) Put(path string, handler Handler) {
	m.handle("PUT", path, handler)
}

func (m *Mux) Delete(path string, handler Handler) {
	m.handle("DELETE", path, handler)
}

func (m *Mux) handle(method, path string, handler Handler) {
	m.entry.Add([]byte(method), []byte(path), handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ps params
	handler := m.entry.Lookup([]byte(r.Method), []byte(r.URL.Path), &ps)
	if handler == nil {
		http.NotFound(w, r)
	} else {
		handler(w, r, ps)
	}
}
