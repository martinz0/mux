package mux

import "net/http"

type Mux interface {
	http.Handler

	Group(path string, f func(m Mux), middleware ...Middleware)

	Get(path string, handler Handler, middleware ...Middleware)
	Post(path string, handler Handler, middleware ...Middleware)
	Put(path string, handler Handler, middleware ...Middleware)
	Delete(path string, handler Handler, middleware ...Middleware)
}

type mux struct {
	entry *muxEntry

	prefix     string
	middleware []Middleware
}

func New() Mux {
	return &mux{
		entry:      NewMuxEntry(),
		middleware: make([]Middleware, 0),
	}
}

func (m *mux) Group(path string, f func(m Mux), middleware ...Middleware) {
	nm := m.Mux(path, middleware...)
	f(nm)
}

func (m *mux) Mux(path string, middleware ...Middleware) Mux {
	nm := m.dup()
	nm.prefix = m.prefix + path
	nm.middleware = append(nm.middleware, m.middleware...)
	nm.middleware = append(nm.middleware, middleware...)
	return nm
}

func (m *mux) dup() *mux {
	nm := new(mux)
	nm.entry = m.entry
	return nm
}

func (m *mux) Get(path string, handler Handler, middleware ...Middleware) {
	m.handle("GET", path, handler, middleware...)
}

func (m *mux) Post(path string, handler Handler, middleware ...Middleware) {
	m.handle("POST", path, handler, middleware...)
}

func (m *mux) Put(path string, handler Handler, middleware ...Middleware) {
	m.handle("PUT", path, handler, middleware...)
}

func (m *mux) Delete(path string, handler Handler, middleware ...Middleware) {
	m.handle("DELETE", path, handler, middleware...)
}

func (m *mux) handle(method, path string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	for i := len(m.middleware) - 1; i >= 0; i-- {
		handler = m.middleware[i](handler)
	}
	path = m.prefix + path
	m.entry.Add(method, []byte(path), handler)
}

func (m *mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	defer ctx.Done()
	m.entry.Lookup(r.Method, []byte(r.URL.Path), ctx.pathParam)(ctx)
}
