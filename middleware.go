package mux

import (
	"log"
)

type Middleware func(Handler) Handler

var (
	PanicRecover = func(next Handler) Handler {
		return func(ctx Context) {
			defer func() {
				if err := recover(); err != nil {
					log.Println(err)
					return
				}
			}()
			next(ctx)
		}
	}
)
