package gpath

import "reflect"

func MapKey(m, k interface{}) (interface{}, bool) {
	if v := MapKeyValue(reflect.ValueOf(m), reflect.ValueOf(k)); v == nil {
		return nil, false
	} else {
		return (*v).Interface(), true
	}
}

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
