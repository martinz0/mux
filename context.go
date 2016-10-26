package mux

import (
	"context"
	"net/http"
)

type Context struct {
	context.Context

	r *http.Request
	w http.ResponseWriter
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Context: context.Background(),
		r: r,
		w: w,
	}
}
