package mux

import (
	"fmt"
	"testing"
	"time"
	"net/http"
)

var m Mux

func init() {
	m = New()
	m.Get("/smartcampus/v1/teachers/classes", TestHandler, Log, PanicRecover)
	http.ListenAndServe(":1234", m)
}

func TestMux(t *testing.T) {
	b := time.Now()
	m := m.(*mux).entry
	m.Lookup("GET", []byte("/smartcampu"))(&Context{})
	m.Lookup("GET", []byte("/smartcampus"))(&Context{})
	m.Lookup("GET", []byte("/smartcampus/v1"))(&Context{})
	m.Lookup("GET", []byte("/smartcampus/v2"))(&Context{})
	m.Lookup("GET", []byte("/smartcampus/v1/teachers/classes"))(&Context{})
	fmt.Println(time.Since(b))
}
