package mux

type Middleware func(next Handler)
