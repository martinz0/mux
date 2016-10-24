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

func (e *muxEntry) addNode(node *muxEntry) {
	if e.nodes == nil {
		e.nodes = make([]*muxEntry, 0)
	}
	e.nodes = append(e.nodes, node)
}

func (e *muxEntry) Lookup(method string, path []byte) Handler {
	if len(path) == 0 || bytes.Equal(path, []byte{'/'}) {
		return NotFoundHandler
	}
	path = bytes.TrimPrefix(path, []byte{'/'})
	path = bytes.TrimSuffix(path, []byte{'/'})

	h := e.lookup(method, path)
	if h == nil {
		return NotFoundHandler
	}
	return h
}

func (e *muxEntry) lookup(method string, path []byte) Handler {
	idx := bytes.Index(path, []byte{'/'})
	if idx == -1 {
		if bytes.Equal(e.part, path) {
			for _, entry := range e.entries {
				if entry.method == method {
					return entry.handler
				}
			}
		}
	} else {
		part := path[:idx]
		if bytes.Equal(e.part, part) {
			for _, node := range e.nodes {
				path = path[idx+1:]
				h := node.lookup(method, path)
				if h == nil {
					continue
				}
				return h
			}
		}
	}
	return nil
}
