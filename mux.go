package mux

import "net/http"

type Mux struct {
	entry *muxEntry

	prefix     string
	middleware []Middleware
}

func New() *Mux {
	return &Mux{
		entry:      NewMuxEntry(),
		middleware: make([]Middleware, 0),
	}
}

func (m *Mux) Group(path string, f func(m *Mux), middleware ...Middleware) {
	nm := m.Mux(path, middleware...)
	f(nm)
}

func (m *Mux) Mux(path string, middleware ...Middleware) *Mux {
	nm := m.dup()
	nm.prefix = m.prefix + path
	nm.middleware = append(nm.middleware, m.middleware...)
	nm.middleware = append(nm.middleware, middleware...)
	return nm
}

func (m *Mux) dup() *Mux {
	nm := new(Mux)
	nm.entry = m.entry
	return nm
}

func (m *Mux) Get(path string, handler Handler, middleware ...Middleware) {
	m.handle("GET", path, handler, middleware...)
}

func (m *Mux) Post(path string, handler Handler, middleware ...Middleware) {
	m.handle("POST", path, handler, middleware...)
}

func (m *Mux) Put(path string, handler Handler, middleware ...Middleware) {
	m.handle("PUT", path, handler, middleware...)
}

func (m *Mux) Delete(path string, handler Handler, middleware ...Middleware) {
	m.handle("DELETE", path, handler, middleware...)
}

func (m *Mux) handle(method, path string, handler Handler, middleware ...Middleware) {
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	for i := len(m.middleware) - 1; i >= 0; i-- {
		handler = m.middleware[i](handler)
	}
	path = m.prefix + path
	m.entry.Add([]byte(method), []byte(path), handler)
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := NewContext(w, r)
	defer ctx.Done()
	m.entry.Lookup([]byte(r.Method), []byte(r.URL.Path), ctx.pathParams)(ctx)
}
