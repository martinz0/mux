package mux

import (
	"bytes"
	"strings"
)

const (
	slash         = "/"
	aliasHolder   = ":"
	aliasPrefix   = ':'
	aliasAsterisk = "*"
)

type muxEntry struct {
	// the front part of path
	part string
	// :alias
	alias string
	// methods and assoicated handlers
	entries []entry
	// trie
	nodes []*muxEntry
}

func (e *muxEntry) String() string {
	var buf bytes.Buffer
	var f func(node *muxEntry)
	f = func(node *muxEntry) {
		for _, n := range node.nodes {
			buf.WriteString(n.part + n.alias)
			buf.WriteByte('\n')
			f(n)
		}
	}
	f(e)
	return buf.String()
}

type entry struct {
	method  string
	handler Handler
}

func (e *muxEntry) setAlias(alias string) {
	if e.alias != "" {
		panic("the muxEntry part alias set")
	}
	e.alias = alias
}

func (e *muxEntry) trimSlash(path string) string {
	path = strings.TrimPrefix(path, slash)
	path = strings.TrimSuffix(path, slash)
	return path
}

func (e *muxEntry) Lookup(method, path string, ps **Params) Handler {
	path = e.trimSlash(path)

	cmu.Lock()
	ca, ok := cache.Get(path)
	cmu.Unlock()

	var me *muxEntry
	if ok {
		*ps = ca.(*Cache).ps
		me = ca.(*Cache).me
	} else {
		me = e.findPath(path, ps)
		cmu.Lock()
		cache.Add(path, &Cache{
			me: me,
			ps: *ps,
		})
		cmu.Unlock()
	}

	if me != nil {
		for _, entry := range me.entries {
			if entry.method == method {
				return entry.handler
			}
		}
	}
	return nil
}

func (e *muxEntry) findPath(path string, ps **Params) *muxEntry {
	if path == "" || e.part == aliasAsterisk {
		return e
	}
	idx := strings.IndexByte(path, '/')
	if idx < 0 {
		return e.find(path, ps)
	}
	me := e.find(path[:idx], ps)
	if me == nil {
		return nil
	}
	return me.findPath(path[idx+1:], ps)
}

func (e *muxEntry) find(path string, ps **Params) *muxEntry {
	holderIdx := -1
	for idx, node := range e.nodes {
		if node.part == path || node.part == aliasAsterisk {
			return node
		}
		if node.part == aliasHolder {
			holderIdx = idx
		}
	}
	if holderIdx > -1 {
		node := e.nodes[holderIdx]
		if *ps == nil {
			p := psPool.Get().(*Params)
			*ps = p
		}
		(*ps).Set(node.alias, strings.TrimSpace(path))
		return node
	}
	return nil
}

func (e *muxEntry) Add(method, path string, handler Handler) {
	path = e.trimSlash(path)
	me := e.add(path)
	for _, entry := range me.entries {
		if entry.method == method {
			panic("muxEntry: add duplicate entry")
		}
	}
	me.entries = append(me.entries, entry{method, handler})
}

func (e *muxEntry) add(path string) *muxEntry {
	if path == "" {
		return e
	}
	idx := strings.IndexByte(path, '/')
	if idx < 0 {
		return e.addPath(path)
	}
	node := e.addPath(path[:idx])
	return node.add(path[idx+1:])
}

func (e *muxEntry) addPath(path string) *muxEntry {
	part := path
	if len(path) > 1 && path[0] == aliasPrefix {
		part = aliasHolder
	}
	node := e.find(part, nil)
	if node == nil {
		node = &muxEntry{
			part:    part,
			entries: make([]entry, 0),
			nodes:   make([]*muxEntry, 0),
		}
		if part == aliasHolder {
			node.setAlias(path[1:])
		}
		e.nodes = append(e.nodes, node)
	}
	return node
}
