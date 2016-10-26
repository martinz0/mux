package mux

import (
	"bytes"
)

type muxEntry struct {
	// the front part of path
	part []byte
	// :alias
	alias []byte
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

func (e *muxEntry) setAlias(alias []byte) {
	if len(e.alias) > 0 {
		panic("the muxEntry part alias set")
	}
	e.alias = alias
}

func (e *muxEntry) trimSlash(path []byte) []byte {
	path = bytes.TrimPrefix(path, []byte{'/'})
	path = bytes.TrimSuffix(path, []byte{'/'})
	return path
}

func (e *muxEntry) Lookup(method string, path []byte, param PathParam) Handler {
	path = e.trimSlash(path)
	h := e.lookup(method, path, param)
	if h == nil {
		h = NotFoundHandler
	}
	return h
}

func (e *muxEntry) lookup(method string, path []byte, param PathParam) Handler {
	me := e.Find(path, param)
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

func (e *muxEntry) Find(path []byte, param PathParam) *muxEntry {
	path = e.trimSlash(path)

	fields := bytes.Split(path, []byte{'/'})
	me := e
	for _, field := range fields {
		if me == nil {
			return nil
		}
		me = me.find(field, param)
	}
	if me == e {
		me = nil
	}
	return me
}

func (e *muxEntry) find(path []byte, param PathParam) *muxEntry {
	for _, node := range e.nodes {
		if bytes.Equal(node.part, path) {
			return node
		}
	}
	if !bytes.Equal(path, []byte("_:_")) {
		for _, node := range e.nodes {
			if bytes.Equal(node.part, []byte("_:_")) {
				if param != nil {
					param[string(node.alias)] = bytes.TrimSpace(path)
				}
				return node
			}
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
		if len(field) > 1 && field[0] == ':' {
			field = []byte("_:_")
		}
		m := me.find(field, nil)
		if m == nil {
			idx--
			break
		}
		m.setAlias(fields[idx][1:])
		me = m
	}
	if idx < len(fields)-1 {
		for _, field := range fields[idx+1:] {
			nm := &muxEntry{
				entries: make([]*entry, 0),
				nodes:   make([]*muxEntry, 0),
			}
			if len(field) > 1 && field[0] == ':' {
				nm.part = []byte("_:_")
				nm.setAlias(field[1:])
			} else {
				nm.part = field
			}
			me.nodes = append(me.nodes, nm)
			me = nm
		}
	}
	return me
}
