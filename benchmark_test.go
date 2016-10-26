package mux

import (
	"testing"
)

var path = []byte("/smartcampus/v1/teachers/10000/classes/100/students")

func Benchmark_Mux(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		nm.entry.Lookup("GET", path, nil)
	}
}
