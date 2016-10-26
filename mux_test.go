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
	m.Group("/smartcampus/v1", func(nm Mux) {
		nm.Get("/teachers/:teacher_id/classes", TestHandler)
		nm.Get("/classes/:class_id/students", TestHandler)
		nm.Get("/students/:student_id/amends", TestHandler)
	}, Log, PanicRecover)

	m.Group("/smartcampus/v1", func(nm Mux) {
		nm.Get("/students/:student_id", TestHandler)
	}, Log)

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
