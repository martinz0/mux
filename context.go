package mux

import (
	"context"
	"net/http"
)

type Context struct {
	context.Context

	r *http.Request
	w http.ResponseWriter

	pathParam PathParam
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Context:   context.Background(),
		r:         r,
		w:         w,
		pathParam: NewPathParam(),
	}
}

func (c *Context) Var(alias string) []byte {
	return c.pathParam[alias]
}

func (c *Context) Write(data []byte) {
	c.w.Write(data)
}
