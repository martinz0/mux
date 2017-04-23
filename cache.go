package mux

import (
	"sync"

	"github.com/golang/groupcache/lru"
)

var (
	cmu   sync.Mutex
	cache = newCache()
)

func newCache() *lru.Cache {
	cache := lru.New(100)
	cache.OnEvicted = func(key lru.Key, value interface{}) {
		ps := value.(*Cache).ps
		if ps != nil {
			ps.reset()
			psPool.Put(ps)
		}
	}
	return cache
}

type Cache struct {
	me *muxEntry
	ps *Params
}
