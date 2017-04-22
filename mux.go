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
	ps := pool.Get().(*Params)
	*ps = (*ps)[:0]
	handler := m.entry.Lookup(r.Method, r.URL.Path, ps)
	if handler == nil {
		http.NotFound(w, r)
	} else {
		handler(w, r, *ps)
	}
	pool.Put(ps)
}

var pool = sync.Pool{
	New: func() interface{} {
		var ps Params
		return &ps
	},
}
