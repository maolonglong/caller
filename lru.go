// Copyright 2022 MaoLongLong. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package caller

import (
	"container/list"
	"runtime"
	"sync"
)

type lruCache struct {
	maxEntries int
	ll         *list.List
	cache      map[uintptr]*list.Element
	mu         sync.Mutex
}

type entry struct {
	key   uintptr
	value runtime.Frame
}

func newLRUCache(maxEntries int) *lruCache {
	return &lruCache{
		maxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[uintptr]*list.Element),
	}
}

func (c *lruCache) store(key uintptr, value runtime.Frame) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		e.Value.(*entry).value = value
		return
	}
	e := c.ll.PushFront(&entry{key, value})
	c.cache[key] = e
	if c.maxEntries != 0 && c.ll.Len() > c.maxEntries {
		c.removeOldestLocked()
	}
}

func (c *lruCache) load(key uintptr) (value runtime.Frame, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return
}

func (c *lruCache) removeOldestLocked() {
	e := c.ll.Back()
	if e != nil {
		c.removeElementLocked(e)
	}
}

func (c *lruCache) removeElementLocked(e *list.Element) {
	c.ll.Remove(e)
	delete(c.cache, e.Value.(*entry).key)
}
