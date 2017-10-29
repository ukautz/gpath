package gpath

import "reflect"

// SliceIndexValue returns value of element with provided index in provided slice. Second return
// value is false, if requested index was out of bounds or index is not integer type or provided
// slice is not actually a slice
func SliceIndex(in interface{}, idx int) (interface{}, bool) {
	if v := SliceIndexValue(reflect.ValueOf(in), idx); v == nil {
		return nil, false
	} else {
		return (*v).Interface(), true
	}
}

// SliceIndexValue returns pointer to reflect.Value of element with provided index in provided slice
// or nil, if requested index was out of bounds or index is not integer type or // provided slice is
// not actually a slice
func SliceIndexValue(r reflect.Value, idx int) *reflect.Value {
	if idx < 0 || r.Kind() != reflect.Slice || idx >= r.Len() {
		return nil
	}
	v := r.Index(idx)
	return &v
}
