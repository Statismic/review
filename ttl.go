package main

import (
	"os"
	"sync"
	"time"
)

type item struct {
	value      *info
	lastAccess int64
}

type ttlMap struct {
	m map[string]*item
	l sync.Mutex
}

func newTTLMap(maxTTL time.Duration) (m *ttlMap) {
	m = &ttlMap{m: make(map[string]*item)}
	go func() {
		for now := range time.Tick(time.Second) {
			m.l.Lock()
			for k, v := range m.m {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					os.RemoveAll(v.value.masterDir)
					os.RemoveAll(v.value.prDir)
					delete(m.m, k)
				}
			}
			m.l.Unlock()
		}
	}()
	return
}

func (m *ttlMap) Len() int {
	return len(m.m)
}

func (m *ttlMap) Put(k string, v *info) {
	m.l.Lock()
	it, ok := m.m[k]
	if !ok {
		it = &item{value: v}
		m.m[k] = it
	}
	it.lastAccess = time.Now().Unix()
	m.l.Unlock()
}

func (m *ttlMap) Get(k string) (v *info) {
	m.l.Lock()
	if it, ok := m.m[k]; ok {
		v = it.value
		it.lastAccess = time.Now().Unix()
	}
	m.l.Unlock()
	return
}
