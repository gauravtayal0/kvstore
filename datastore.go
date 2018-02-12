package kvm

import (
	"sync"
)

type KV struct {
	Key   string
	Value string
}

type DataStore struct {
	m   map[string]KV
	mux sync.RWMutex
}

func New() DataStore {
	ds := DataStore{m: make(map[string]KV)}
	return ds
}

func (ds *DataStore) Set(key, value string) {
	ds.mux.Lock()
	ds.m[key] = KV{key, value}
	ds.mux.Unlock()
}

func (ds *DataStore) Get(key string) (KV, bool) {
	ds.mux.RLock()
	kvpair, ok := ds.m[key]
	ds.mux.RUnlock()

	return kvpair, ok
}

func (ds *DataStore) GetValue(key string) (string, bool) {
	ds.mux.RLock()
	kv, ok := ds.m[key]
	ds.mux.RUnlock()

	return kv.Value, ok
}

func (ds *DataStore) Delete(key string) {
	delete(ds.m, key)
}

// Exists checks for the existence of key in the store.
func (ds *DataStore) Exists(key string) bool {
	_, ok := ds.Get(key)

	if ok {
		return true
	}
	return false
}
