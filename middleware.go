package mux

import (
	"log"
	"time"
)

type Middleware func(Handler) Handler

var (
	PanicRecover = func(next Handler) Handler {
		return func(ctx *Context) {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
					return
				}
			}()
			next(ctx)
		}
	}

	Log = func(next Handler) Handler {
		return func(ctx *Context) {
			b := time.Now()
			next(ctx)
			log.Println(time.Since(b))
		}
	}
)
