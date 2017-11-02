// gpath is a library for
package gpath

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"errors"
)

// GPath provides path based access to map or slice data structures in Go. Additionally type conversion
// helper methods are provided, helping to work with user input.
type GPath struct {
	source     interface{}
	traversals *cache
}

var vof = reflect.ValueOf

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
		r := vof(val)
		return r.Kind() == reflect.Slice
	}
	return false
}

// IsMap returns bool whether path exists AND is a map of some kind
func (gp *GPath) IsMap(path string) bool {
	if val, has := gp.get(path); has {
		r := vof(val)
		return r.Kind() == reflect.Map
	}
	return false
}

// Get returns the value of path - or nil, if it does not exist
func (gp *GPath) Get(path string) interface{} {
	val, _ := gp.get(path)
	return val
}

// Set creates or writes a new value with given path. Only child elements can be modified.
func (gp *GPath) Set(path string, value interface{}) error {
	var to interface{}
	key := ""
	root := ""
	if idx := strings.LastIndex(path, "."); idx <= 0 {
		to = gp.source
		key = path
		root = "."
	} else if parent := gp.Get(path[0:idx]); parent != nil {
		to = parent
		key = path[idx+1:]
		root = "." + path[0:idx]
	} else {
		return fmt.Errorf("parent element %s does not exist", path[0:idx])
	}

	set := false
	ref := vof(to)
	refk := ref.Kind()
	isref := true
	if refk == reflect.Slice || refk == reflect.Map {
		ptr := reflect.New(ref.Type())
		ptr.Elem().Set(ref)
		ref = ptr
		refk = ref.Kind()
		isref = false
	}
	if refk == reflect.Ptr {
		switch ref.Elem().Kind() {
		case reflect.Slice:
			if key != "-1" && !isUInt(key) {
				return fmt.Errorf("cannot write to \"%s\" as %s is a slice, not a map, and requires positive integer indices", key, root)
			} else if root == "." && !isref {
				return errors.New("parent element cannot be slice. Provide either pointer to slice or slice embedded within maps")
			}
			idx, _ := strconv.Atoi(key)
			if set = SliceIndexSet(ref.Interface(), idx, value, true); set {
				if !isref && root != "." {
					return gp.Set(root[1:], ref.Elem().Interface())
				} else {
					return nil
				}
			}
		case reflect.Map:
			set = MapKeySet(ref.Elem().Interface(), key, value, true)
		}
	}
	if set {
		gp.traversals.set(path, value)
		gp.traversals.clear(path + ".")
		return nil
	} else {
		return fmt.Errorf("could not set %s in %s (%s)", path, root, reflect.ValueOf(to).Kind())
	}

}

func (gp *GPath) get(path string) (interface{}, bool) {
	if val, has := gp.traversals.get(path); has {
		return val, true
	} else if val, has = followPath(path, gp.source); has {
		gp.traversals.set(path, val)
		return val, true
	} else {
		gp.traversals.set(path, nil)
		return nil, false
	}
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
	return p[0], p[1:]
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

func (c *cache) clear(prefix string) int {
	count := 0
	keys := []string{}
	c.mux.Lock()
	defer c.mux.Unlock()
	for key, _ := range c.data {
		keys = append(keys, key)
	}
	for _, key := range keys {
		if strings.HasPrefix(key, prefix) {
			count++
			delete(c.data, key)
		}
	}
	return count
}
