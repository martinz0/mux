package mux

type Handler func(*Context)

var NotFoundHandler Handler = func(ctx *Context) {
	ctx.w.Write([]byte("not found"))
}

var TestHandler Handler = func(ctx *Context) {
	for key, val := range ctx.pathParam {
		ctx.w.Write([]byte(key))
		ctx.w.Write([]byte{':'})
		ctx.w.Write(val)
		ctx.w.Write([]byte{'\n'})
	}
}
