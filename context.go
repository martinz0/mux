package mux

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

type Context struct {
	context.Context

	r *http.Request
	w http.ResponseWriter

	pathParam  PathParam
	queryParam QueryParam
	body       []byte

	response []byte
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	r.ParseForm()
	queryParam := NewQueryParam()
	for key, _ := range r.Form {
		queryParam[strings.TrimSpace(key)] = strings.TrimSpace(r.Form.Get(key))
	}

	var data []byte
	if r.Method == "POST" || r.Method == "PUT" {
		data, _ = ioutil.ReadAll(r.Body)
	}

	return &Context{
		Context:    context.Background(),
		r:          r,
		w:          w,
		pathParam:  NewPathParam(),
		queryParam: queryParam,
		body:       data,
	}
}

func (c *Context) Var(alias string) []byte {
	return c.pathParam[alias]
}

func (c *Context) Write(data []byte) {
	c.response = data
}

func (c *Context) Done() {
	c.w.Write(c.response)
}
