package mux

type Handler func(Context)

var NotFoundHandler Handler = func(ctx Context) {
}

var TestHandler Handler = func(ctx Context) {
}
