package mux

import "net/http"

type Mux interface {
	http.Handler

	// FIXME
	Group(path string, middleware ...Middleware) Mux
	Get(path string, handler Handler, middleware ...Middleware)
}

type mux struct {
	entry *muxEntry
}

func New() Mux {
	return &mux{
		entry: NewMuxEntry(),
	}
}

func (m *mux) dup() *mux {
	nm := new(mux)
	nm.entry = new(muxEntry)
	*nm.entry = *m.entry
	return nm
}

func (m *mux) Group(path string, middleware ...Middleware) Mux {
	nm := m.dup()
	nm.entry.groupPrefix = path
	nm.entry.groupMiddleware = middleware
	return nm
}

func (m *mux) Get(path string, handler Handler, middleware ...Middleware) {
	m.handle("GET", path, handler, middleware...)
}

func (m *mux) handle(method, path string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	for i := len(m.entry.groupMiddleware) - 1; i >= 0; i-- {
		handler = m.entry.groupMiddleware[i](handler)
	}
	path += m.entry.groupPrefix
	m.entry.Add(method, []byte(path), handler)
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	m.entry.Lookup(r.Method, []byte(r.URL.Path), ctx.pathParam)(ctx)
}
