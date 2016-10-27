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

	response []byte
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
		Context:     context.Background(),
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

func (c *Context) Done() {
	c.w.Write(c.response)
}
