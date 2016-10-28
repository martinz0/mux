package mux

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

type Context struct {
	context.Context

	r *http.Request
	w http.ResponseWriter

	pathParams  *params
	queryParams *params
	body        []byte

	statusCode int
	response   []byte
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	r.ParseForm()
	query := NewParams()
	for key, _ := range r.Form {
		query.Set(bytes.TrimSpace([]byte(key)), bytes.TrimSpace([]byte(r.Form.Get(key))))
	}

	var data []byte
	if r.Method == "POST" || r.Method == "PUT" {
		data, _ = ioutil.ReadAll(r.Body)
	}

	return &Context{
		Context:     r.Context(),
		r:           r,
		w:           w,
		pathParams:  NewParams(),
		queryParams: query,
		body:        data,
	}
}

func (c *Context) Write(data []byte) {
	c.response = data
}

func (c *Context) WriteHeader(statusCode int) {
	c.statusCode = statusCode
}

func (c *Context) Done() {
	if c.w != nil {
		if c.statusCode > 0 {
			c.w.WriteHeader(c.statusCode)
		}
		if len(c.response) > 0 {
			c.w.Write(c.response)
		}
	}
}
