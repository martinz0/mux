package mux

type Handler func(*Context)

var NotFoundHandler Handler = func(ctx *Context) {
	ctx.w.Write([]byte("not found"))
}

var TestHandler Handler = func(ctx *Context) {
	ctx.Write(ctx.body)
}
