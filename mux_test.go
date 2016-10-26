package mux

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var m Mux

func init() {
	m = New()
	// m.Get("/smartcampus/v1/teachers/:teacher_id/classes/:class_id/students/:student_id", TestHandler, Log, PanicRecover)

	g := m.Group("/smartcampus/v1", Log, PanicRecover)
	g.Get("/teachers/:teacher_id/classes", TestHandler)

	http.ListenAndServe(":1234", m)
}

func TestMux(t *testing.T) {
	b := time.Now()
	param := NewPathParam()
	m := m.(*mux).entry
	/*
		m.Lookup("GET", []byte("/smartcampu"))(&Context{})
		m.Lookup("GET", []byte("/smartcampus"))(&Context{})
		m.Lookup("GET", []byte("/smartcampus/v1"))(&Context{})
		m.Lookup("GET", []byte("/smartcampus/v2"))(&Context{})
		m.Lookup("GET", []byte("/smartcampus/v1/teachers/classes"))(&Context{})
	*/
	m.Lookup("GET", []byte("/smartcampus/v1/teachers/10000/classes"), param)(&Context{})
	fmt.Println(time.Since(b))
}
