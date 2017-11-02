package gpath

import (
	"fmt"
	"github.com/ukautz/cast"
	"reflect"
)

// SliceIndexValue returns value of element with provided index in provided slice. Second return
// value is false, if requested index was out of bounds or index is not integer type or provided
// slice is not actually a slice
func SliceIndex(theSlice interface{}, idx int) (interface{}, bool) {
	if v := SliceIndexValue(vof(theSlice), idx); v == nil {
		return nil, false
	} else {
		return (*v).Interface(), true
	}
}
func SliceIndexSet(theSlice interface{}, idx int, theValue interface{}, castFitting ...bool) bool {
	return SliceIndexValueSet(vof(theSlice), idx, vof(theValue), castFitting...) == nil
}

// SliceIndexValue returns pointer to reflect.Value of element with provided index in provided slice
// or nil, if requested index was out of bounds or index is not integer type or // provided slice is
// not actually a slice
func SliceIndexValue(theSlice reflect.Value, idx int) *reflect.Value {
	if idx < 0 || theSlice.Kind() != reflect.Slice || idx >= theSlice.Len() {
		return nil
	}
	v := theSlice.Index(idx)
	return &v
}

// SliceIndexValueSet
func SliceIndexValueSet(theSlice reflect.Value, idx int, theValue reflect.Value, castFitting ...bool) error {
	actual := reflect.Indirect(theSlice)
	if theSlice.Kind() != reflect.Ptr || actual.Kind() != reflect.Slice {
		fmt.Printf("Cannot set %s into %s (%s aka %s) slice\n", theValue.Kind(), theSlice.Kind(), actual.Kind(), reflect.TypeOf(actual).Kind())
		return fmt.Errorf("expected pointer to Slice but got %s (indirect: %s)", theSlice.Kind(), actual.Kind())
	} else if idx < -1 || idx > actual.Len() {
		return fmt.Errorf("provided index %d is out of bounds for slice of len %d", idx, actual.Len())
	}

	ek := actual.Type().Elem().Kind()
	if ek != reflect.Interface {
		if vk := theValue.Kind(); vk != ek {
			if len(castFitting) > 0 && castFitting[0] {
				if ref := cast.CastToValue(theValue, ek); ref != nil {
					theValue = *ref
				} else {
					return fmt.Errorf("provided value is of %s kind and cannot be cast into %s kind", theValue.Kind(), ek)
				}
			} else {
				return fmt.Errorf("provided value is of %s kind but must be %s kind", theValue.Kind(), ek)
			}
		}
	}

	if idx == -1 || idx == actual.Len() {
		actual.Set(reflect.Append(actual, theValue))
	} else {
		actual.Index(idx).Set(theValue)
	}
	return nil
}
