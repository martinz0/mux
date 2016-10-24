package mux

import (
	"fmt"
	"time"
	"testing"
)

var me *muxEntry
func init() {
	me = &muxEntry{
		part: []byte("smartcampus"),
		entries: []*entry{
			{"GET", func(Context){println("GET /smartcampus")}},
		},
	}
	me.addNode(&muxEntry{
		part: []byte("v1"),
		entries: []*entry{
			{"GET", func(Context){println("GET /smartcampus/v1")}},
		},
	})
}

func TestMux(t *testing.T) {
	b := time.Now()
	h := me.Lookup("GET", []byte("/smartcampu"))
	h(Context{})
	h = me.Lookup("GET", []byte("/smartcampus"))
	h(Context{})
	h = me.Lookup("GET", []byte("/smartcampus/v1"))
	h(Context{})
	h = me.Lookup("GET", []byte("/smartcampus/v2"))
	h(Context{})
	fmt.Println(time.Since(b))
}
