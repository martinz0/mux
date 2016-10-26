package mux

import (
	"fmt"
	"testing"
	"time"
)

var m *muxEntry

func init() {
	m = NewMuxEntry()
	m.Add("GET", []byte("smartcampus/v1/teachers/classes"), TestHandler)
}

func TestMux(t *testing.T) {
	b := time.Now()
	m.Lookup("GET", []byte("/smartcampu"))(Context{})
	m.Lookup("GET", []byte("/smartcampus"))(Context{})
	m.Lookup("GET", []byte("/smartcampus/v1"))(Context{})
	m.Lookup("GET", []byte("/smartcampus/v2"))(Context{})
	m.Lookup("GET", []byte("/smartcampus/v1/teachers/classes"))(Context{})
	fmt.Println(time.Since(b))
}
