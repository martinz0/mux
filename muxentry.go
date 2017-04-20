package mux

import (
	"bytes"
	"net/http"
)

var (
	slash            = []byte{'/'}
	aliasHolder      = []byte("_:_")
	aliasPrefix byte = ':'
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
	method  []byte
	handler http.Handler
}

func NewMuxEntry() *muxEntry {
	return &muxEntry{
		entries: make([]*entry, 0),
		nodes:   make([]*muxEntry, 0),
	}
}

func (e *muxEntry) setAlias(alias []byte) {
	if len(e.alias) == 0 {
		e.alias = alias
	}
	if len(e.alias) > 0 && !bytes.Equal(e.alias, alias) {
		panic("the muxEntry part alias set")
	}
}

func (e *muxEntry) trimSlash(path []byte) []byte {
	path = bytes.TrimPrefix(path, slash)
	path = bytes.TrimSuffix(path, slash)
	return path
}

func (e *muxEntry) Lookup(method, path []byte, p *params) http.Handler {
	path = e.trimSlash(path)
	h := e.lookup(method, path, p)
	if h == nil {
		h = http.NotFoundHandler()
	}
	return h
}

func (e *muxEntry) lookup(method, path []byte, p *params) http.Handler {
	me := e.findPath(path, p)
	if me == nil {
		return nil
	}
	for _, entry := range me.entries {
		if bytes.Equal(entry.method, method) {
			return entry.handler
		}
	}
	return nil
}

func (e *muxEntry) findPath(path []byte, p *params) *muxEntry {
	me := e

	var idx int
	for idx > -1 {
		if me == nil {
			return nil
		}
		idx = bytes.IndexByte(path, '/')
		if idx > 0 {
			me = me.find(path[:idx], p)
			path = path[idx+1:]
		} else {
			me = me.find(path, p)
		}
	}
	if me == e {
		me = nil
	}
	return me
}

func (e *muxEntry) find(path []byte, p *params) *muxEntry {
	for _, node := range e.nodes {
		if bytes.Equal(node.part, path) {
			return node
		}
	}
	if !bytes.Equal(path, aliasHolder) {
		for _, node := range e.nodes {
			if bytes.Equal(node.part, aliasHolder) {
				p.Set(node.alias, bytes.TrimSpace(path))
				return node
			}
		}
	}
	return nil
}

func (e *muxEntry) Add(method, path []byte, handler http.Handler) {
	path = e.trimSlash(path)
	me := e.add(path)
	for _, entry := range me.entries {
		if bytes.Equal(entry.method, method) {
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
		fields = bytes.Split(path, slash)
	)
	for idx, field = range fields {
		if len(field) > 1 && field[0] == aliasPrefix {
			field = aliasHolder
		}
		m := me.find(field, nil)
		if m == nil {
			idx--
			break
		}
		if bytes.Equal(field, aliasHolder) {
			m.setAlias(fields[idx][1:])
		}
		me = m
	}
	if idx < len(fields)-1 {
		for _, field := range fields[idx+1:] {
			nm := &muxEntry{
				entries: make([]*entry, 0),
				nodes:   make([]*muxEntry, 0),
			}
			if len(field) > 1 && field[0] == aliasPrefix {
				nm.part = aliasHolder
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
