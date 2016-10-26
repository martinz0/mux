package mux

import (
	"bytes"
)

type muxEntry struct {
	// the front part of path
	part []byte
	// methods and assoicated handlers
	entries []*entry
	// trie
	nodes []*muxEntry
}

type entry struct {
	method  string
	handler Handler
}

func NewMuxEntry() *muxEntry {
	return &muxEntry{
		part:    make([]byte, 0),
		entries: make([]*entry, 0),
		nodes:   make([]*muxEntry, 0),
	}
}

func (e *muxEntry) trimSlash(path []byte) []byte {
	path = bytes.TrimPrefix(path, []byte{'/'})
	path = bytes.TrimSuffix(path, []byte{'/'})
	return path
}

func (e *muxEntry) Lookup(method string, path []byte) Handler {
	path = e.trimSlash(path)
	h := e.lookup(method, path)
	if h == nil {
		h = NotFoundHandler
	}
	return h
}

func (e *muxEntry) lookup(method string, path []byte) Handler {
	me := e.Find(path)
	if me == nil {
		return nil
	}
	for _, entry := range me.entries {
		if entry.method == method {
			return entry.handler
		}
	}
	return nil
}

func (e *muxEntry) Find(path []byte) *muxEntry {
	path = e.trimSlash(path)

	fields := bytes.Split(path, []byte{'/'})
	me := e
	for _, field := range fields {
		if me == nil {
			return nil
		}
		me = me.find(field)
	}
	if me == e {
		me = nil
	}
	return me
}

func (e *muxEntry) find(path []byte) *muxEntry {
	for _, node := range e.nodes {
		if bytes.Equal(node.part, path) {
			return node
		}
	}
	return nil
}

func (e *muxEntry) Add(method string, path []byte, handler Handler) {
	path = e.trimSlash(path)
	me := e.add(path)
	for _, entry := range me.entries {
		if entry.method == method {
			panic("muxEntry: add duplicate entry")
		}
	}
	me.entries = append(me.entries, &entry{method, handler})
}

func (e *muxEntry) add(path []byte) *muxEntry {
	var (
		me     = e
		idx    int
		field  []byte
		fields = bytes.Split(path, []byte{'/'})
	)
	for idx, field = range fields {
		m := me.find(field)
		if m == nil {
			idx--
			break
		}
		me = m
	}
	if idx < len(fields)-1 {
		for _, field := range fields[idx+1:] {
			nm := &muxEntry{
				part:    field,
				entries: make([]*entry, 0),
				nodes:   make([]*muxEntry, 0),
			}
			me.nodes = append(me.nodes, nm)
			me = nm
		}
	}
	return me
}
