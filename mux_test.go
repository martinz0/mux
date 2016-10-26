package mux

import (
// "net/http"
)

var m Mux
var nm *mux

func init() {
	m = New()
	nm = m.(*mux)
	m.Get("/smartcampus/v1/teachers/:teacher_id/classes/:class_id/students", TestHandler)
	/*
		m.Group("/smartcampus/v1", func(nm Mux) {
			nm.Get("/teachers/:teacher_id/classes", TestHandler)
			nm.Get("/classes/:class_id/students", TestHandler)
			nm.Get("/students/:student_id/amends", TestHandler)
		}, Log, PanicRecover)

		m.Group("/smartcampus/v1", func(nm Mux) {
			nm.Get("/students/:student_id", TestHandler)
		}, Log)

		http.ListenAndServe(":1234", m)
	*/
}
