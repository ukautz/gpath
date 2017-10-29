package gpath

import "reflect"

func SliceIndex(in interface{}, idx int) (interface{}, bool) {
	if v := SliceIndexValue(reflect.ValueOf(in), idx); v == nil {
		return nil, false
	} else {
		return (*v).Interface(), true
	}
}

func SliceIndexValue(r reflect.Value, idx int) *reflect.Value {
	if idx < 0 || r.Kind() != reflect.Slice || r.Len() <= idx {
		return nil
	}
	v := r.Index(idx)
	return &v
}
