package mux

type Mux interface {
	Get(path string, handler Handler, middleware ...Middleware)
}

type mux struct {
	entry *muxEntry
}

func New() Mux {
	return new(mux)
}

func (m *mux) Get(path string, handler Handler, middleware ...Middleware) {
}
