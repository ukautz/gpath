// gpath is a library for
package gpath

import (
	"reflect"
	"strconv"
	"strings"
	"sync"
)

// GPath provides path based access to map or slice data structures in Go. Additionally type conversion
// helper methods are provided, helping to work with user input.
type GPath struct {
	source     interface{}
	traversals *cache
}

// New creates new GPath instance for arbitrary map or slice instances
func New(from interface{}) *GPath {
	return &GPath{
		source:     from,
		traversals: newCache(map[string]interface{}{}),
	}
}

// Has returns bool whether given path exists
func (gp *GPath) Has(path string) bool {
	_, has := gp.get(path)
	return has
}

// IsSlice returns bool whether path exists AND is a slice of some kind
func (gp *GPath) IsSlice(path string) bool {
	if val, has := gp.get(path); has {
		r := reflect.ValueOf(val)
		return r.Kind() == reflect.Slice
	}
	return false
}

// IsMap returns bool whether path exists AND is a map of some kind
func (gp *GPath) IsMap(path string) bool {
	if val, has := gp.get(path); has {
		r := reflect.ValueOf(val)
		return r.Kind() == reflect.Map
	}
	return false
}

// Get returns the value of path - or nil, if it does not exist
func (gp *GPath) Get(path string) interface{} {
	val, _ := gp.get(path)
	return val
}

func (gp *GPath) get(path string) (interface{}, bool) {
	if val, has := gp.traversals.get(path); has {
		return val, true
	} else if val, has = followPath(path, gp.source); has {
		gp.traversals.set(path, val)
		return val, true
	} else {
		gp.traversals.set(path, nil)
	}
	return nil, false
}

func getNext(idx string, in interface{}) (interface{}, bool) {
	if isUInt(idx) {
		i, _ := strconv.Atoi(idx)
		return SliceIndex(in, i)
	}
	return MapKey(in, idx)
}

func followPath(path string, in interface{}) (res interface{}, found bool) {
	cur, next := splitPath(path)
	for res, found = getNext(cur, in); found; res, found = getNext(cur, in) {
		if len(next) == 0 {
			return
		}
		in = res
		cur = next[0]
		next = next[1:]
	}
	return
}

func isUInt(word string) bool {
	for _, c := range word {
		if c < '0' || c > '9' {
			return false
		}
	}
	return word != ""
}

func splitPath(path string) (string, []string) {
	p := strings.Split(path, ".")
	if l := len(p); l == 0 {
		return "", []string{}
	} else {
		return p[0], p[1:]
	}
}

// traversals implements concurrency safe map (I want to support Go < 1.9, so no sync.Map ..)
type cache struct {
	data map[string]interface{}
	mux  *sync.RWMutex
}

func newCache(data map[string]interface{}) *cache {
	return &cache{
		data: data,
		mux:  new(sync.RWMutex),
	}
}

func (c *cache) get(key string) (interface{}, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	if v, ok := c.data[key]; ok {
		return v, true
	}
	return nil, false
}

func (c *cache) set(key string, value interface{}) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.data[key] = value
	return value
}
