// gpath is a library for
package gpath

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

	// find parent:
	// path is either of root (no "."), which makes root the parent, or or below root (with "."), which
	// makes the "path's parent" the parent. So path="foo" -> root is parent and "foo.bar" -> "foo" is parent
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

	// make parent a pointer:
	// we want to support *map, map, *slice and slice alike. For simplification,
	// just cast now map -> *map or slice -> *slice
	if refk == reflect.Slice || refk == reflect.Map {
		ptr := reflect.New(ref.Type())
		ptr.Elem().Set(ref)
		ref = ptr
		refk = ref.Kind()
		isref = false
	}

	// at this point, must be pointer || fail
	if refk != reflect.Ptr {
		return fmt.Errorf("could not set %s in %s (%s)", path, root, reflect.ValueOf(to).Kind())
	}

	// now:
	// for slice -> set new value in pointer to slice
	// for map -> just put it in there
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

	// when anything was actually changed -> write in cache + clear all below cache, as it is gone
	if set {
		gp.traversals.set(path, value)
		gp.traversals.clear(path + ".")
		return nil
	} else {
		return fmt.Errorf("could not set %s in %s (%s)", path, root, reflect.ValueOf(to).Kind())
	}
}

// GetChild returns path value as *gpath.GPath (child) object, if the path value is either a Map or a Slice of any kind.
// In case of slice, a reference to the slice is used. Otherwise nil is returned.
func (gp *GPath) GetChild(path string) *GPath {
	if val, ok := gp.get(path); ok {
		ref := reflect.ValueOf(val)
		switch refk := ref.Kind(); refk {
		case reflect.Slice:
			ptr := reflect.New(ref.Type())
			ptr.Elem().Set(ref)
			return New(ptr.Interface())
		case reflect.Map:
			return New(ref.Interface())
		case reflect.Ptr:
			elem := ref.Elem()
			switch elemk := elem.Kind(); elemk {
			case reflect.Slice:
				return New(ref.Interface())
			case reflect.Map:
				return New(ref.Interface())
			}
		}
	}
	return nil
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
