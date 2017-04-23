package mux

import (
	"net/http"
	"sync"
)

type Mux struct {
	entry *muxEntry
}

func New() *Mux {
	return &Mux{
		entry: &muxEntry{
			entries: make([]entry, 0),
			nodes:   make([]*muxEntry, 0),
		},
	}
}

type Handler func(w http.ResponseWriter, r *http.Request, ps Params)

func (m *Mux) Handle(method, path string, handler Handler) {
	m.entry.Add(method, path, handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var ps *Params
	handler := m.entry.Lookup(r.Method, r.URL.Path, &ps)
	switch {
	case handler == nil:
		http.NotFound(w, r)
	case ps == nil:
		handler(w, r, nil)
	default:
		handler(w, r, *ps)
	}
}

var psPool = sync.Pool{
	New: func() interface{} {
		return new(Params)
	},
}
