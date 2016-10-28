package mux

import "net/http"

type Handler func(*Context)

var NotFoundHandler Handler = func(ctx *Context) {
	ctx.WriteHeader(http.StatusNotFound)
}

var TestHandler Handler = func(ctx *Context) {
	ctx.Write(ctx.body)
}
