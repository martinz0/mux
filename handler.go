package mux

type Handler func(*Context)

var NotFoundHandler Handler = func(ctx *Context) {
	ctx.w.Write([]byte("NotFound"))
}

var TestHandler Handler = func(ctx *Context) {
	ctx.w.Write([]byte("Test"))
}
