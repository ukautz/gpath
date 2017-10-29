package gpath

import "reflect"

// MapKey returns value in provided map with given key. Second return parameter is false, if requested
// key was not found (or key type is not compatible with type of keys of map or provided map is not
// actually a map)
func MapKey(m, k interface{}) (interface{}, bool) {
	if v := MapKeyValue(reflect.ValueOf(m), reflect.ValueOf(k)); v == nil {
		return nil, false
	} else {
		return (*v).Interface(), true
	}
}

// MapKeyValue returns pointer to reflect.Value in provided map with given key or nil, if requested
// key was not found (or key type is not compatible with type of keys of map or provided map is not
// actually a map)
func MapKeyValue(m, k reflect.Value) *reflect.Value {
	if m.Kind() != reflect.Map {
		return nil
	}
	kk := m.MapKeys()
	if len(kk) == 0 || kk[0].Kind() != k.Kind() {
		return nil
	} else if v := m.MapIndex(k); !v.IsValid() {
		return nil
	} else {
		return &v
	}
}
