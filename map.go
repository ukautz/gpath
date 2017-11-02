package gpath

import (
	"fmt"
	"reflect"
	"github.com/ukautz/cast"
)

// MapKey returns value in provided map with given key. Second return parameter is false, if requested
// key was not found (or key type is not compatible with type of keys of map or provided map is not
// actually a map)
func MapKey(m, k interface{}) (interface{}, bool) {
	if v := MapKeyValue(vof(m), vof(k)); v == nil {
		return nil, false
	} else {
		return (*v).Interface(), true
	}
}

// MapKeySet inserts given value at given key in given map. Returns bool whether key could be set.
// The 4th optional bool parameter enables type casting, so that
//
//		m := map[string]int8{}
//
//		MapKeySet(m, "key", int8(123)) // returns true
//		// now m has key string("key") with value int(123)
//
//		MapKeySet(m, 123, "33.5") // returns false
//		// m has not changed, since int(123) is not a valid key (and string("33.5") not a valid value)
//
//		MapKeySet(m, 123, "33.5", true) // returns true
//		// now m has casted key string("123") with casted value int8(33)
func MapKeySet(theMap, theKey, theValue interface{}, castFitting ...bool) bool {
	return MapKeyValueSet(vof(theMap), vof(theKey), vof(theValue), castFitting...) == nil
}

// MapKeyValue returns pointer to reflect.Value in provided map with given key or nil, if requested
// key was not found (or key type is not compatible with type of keys of map or provided map is not
// actually a map)
func MapKeyValue(m, k reflect.Value) *reflect.Value {
	if m.Kind() != reflect.Map {
		return nil
	}
	kk := m.MapKeys()
	kki := m.Type().Key().Kind()
	if len(kk) == 0 || (kki != reflect.Interface && kki != k.Kind()) {
		return nil
	} else if v := m.MapIndex(k); !v.IsValid() {
		return nil
	} else {
		return &v
	}
}

// MapKeyValueSet works as MapKeyValue, but it expects reflect.Value parameters and returns specific error
// why key, value could not be added.
func MapKeyValueSet(theMap, theKey, theValue reflect.Value, castFitting ...bool) error {
	if theMap.Kind() != reflect.Map {
		return fmt.Errorf("provided map is not Map kind but %s kind", theMap.Kind())
	}
	t := theMap.Type()
	kk := t.Key().Kind()
	fit := len(castFitting) > 0 && castFitting[0]
	if kk != reflect.Interface && kk != theKey.Kind() {
		if fit {
			if ref := cast.CastToValue(theKey, kk); ref != nil {
				theKey = *ref
			} else {
				return fmt.Errorf("provided value is of %s kind and cannot be cast into %s kind", theKey.Kind(), kk)
			}
		} else {
			return fmt.Errorf("provided key is of %s kind but must be %s kind", theKey.Kind(), kk)
		}
	}
	ek := t.Elem().Kind()
	if ek != reflect.Interface && ek != theValue.Kind() {
		if fit {
			if ref := cast.CastToValue(theValue, ek); ref != nil {
				theValue = *ref
			} else {
				return fmt.Errorf("provided value is of %s kind and cannot be cast into %s kind", theValue.Kind(), ek)
			}
		} else {
			return fmt.Errorf("provided value is of %s kind but must be %s kind", theValue.Kind(), ek)
		}
	}
	theMap.SetMapIndex(theKey, theValue)
	return nil
}
