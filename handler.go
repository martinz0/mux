package mux

type Handler func(Context)

var NotFoundHandler Handler = func(ctx Context) {
	println("not found")
}
