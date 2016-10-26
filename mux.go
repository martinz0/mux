package mux

type Mux interface {
	Get(path string, handler Handler, middleware ...Middleware)
}

type mux struct {
	group      *muxEntry
	entry      *muxEntry
	middleware []Middleware
}

func New() Mux {
	return new(mux)
}

/*
func (m *mux) Group(path string, f func(), middleware ...Middleware) {
	m.group = m.entry.lookup(path)
}
*/

func (m *mux) Get(path string, handler Handler, middleware ...Middleware) {
	m.handle("GET", path, handler, middleware...)
}

func (m *mux) handle(method, path string, handler Handler, middleware ...Middleware) {
}
