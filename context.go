package mux

import (
	"context"
	"net/http"
)

type Context struct {
	context.Context

	http.Request
	http.ResponseWriter
}
