package mux

import (
	"net/http"
	"strings"
	"sync"
)

type Mux struct {
	entry    *muxEntry
	notFound Handler
}

func New() *Mux {
	return &Mux{
		entry: &muxEntry{
			entries: make([]entry, 0),
			nodes:   make([]*muxEntry, 0),
		},
		notFound: notFoundHandler,
	}
}

type Handler func(w http.ResponseWriter, r *http.Request, ps Params)

var notFoundHandler = func(w http.ResponseWriter, r *http.Request, ps Params) {
	http.NotFound(w, r)
}

func (m *Mux) Handle(method, path string, handler Handler) {
	m.entry.Add(method, path, handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ps *Params
	handler := m.entry.Lookup(r.Method, r.URL.Path, &ps)
	if handler == nil {
		if strings.HasPrefix(r.URL.Path, "/debug/pprof/") {
			// support net/http/pprof
			http.DefaultServeMux.ServeHTTP(w, r)
			return
		}
		handler = m.notFound
	}
	if ps == nil {
		handler(w, r, nil)
	} else {
		handler(w, r, *ps)
	}
}

func (m *Mux) SetNotFound(h Handler) {
	m.notFound = h
}

var psPool = sync.Pool{
	New: func() interface{} {
		return new(Params)
	},
}
