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

type Handler func(w http.ResponseWriter, r *http.Request, ps Params)

func (m *Mux) Handle(method, path string, handler Handler) {
	m.entry.Add(method, path, handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ps Params
	handler := m.entry.Lookup(r.Method, r.URL.Path, &ps)
	if handler == nil {
		http.NotFound(w, r)
	} else {
		handler(w, r, ps)
	}
}
