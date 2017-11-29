package gpath

import (
	"github.com/ukautz/cast"
	"reflect"
)

// IsMap returns bool whether path exists AND is a map of some kind
func (gp *GPath) IsMap(path string) bool {
	if val, has := gp.get(path); has {
		r := vof(val)
		return r.Kind() == reflect.Map
	}
	return false
}

// GetMap returns the value of the path as map[interface{}]interface{}, or nil if value of path is not a map
func (gp *GPath) GetMap(path string) map[interface{}]interface{} {
	if val, has := gp.get(path); has {
		if mval, ok := cast.CastMap(val); ok {
			return mval
		}
	}
	return nil
}

// GetMapString returns the value of the path as map[string]interface{}, or nil if value of path is not a map or
// if any map keys are not castable to string
func (gp *GPath) GetMapString(path string, fallback ...map[string]interface{}) map[string]interface{} {
	if val, has := gp.get(path); has {
		if mval, ok := cast.CastMapString(val); ok {
			return mval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return nil
}

// GetMapStringString returns the value of the path as map[string]string, or nil if value of path is not a map or
// if any map keys or values are not castable to string
func (gp *GPath) GetMapStringString(path string, fallback ...map[string]string) map[string]string {
	if val, has := gp.get(path); has {
		if mval, ok := cast.CastMapStringString(val); ok {
			return mval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return nil
}

// GetMapStringInt returns the value of the path as map[string]int64, or nil if value of path is not a map or
// if any map keys are not castable to string or any values not castable to int64
func (gp *GPath) GetMapStringInt(path string, fallback ...map[string]int64) map[string]int64 {
	if val, has := gp.get(path); has {
		if mval, ok := cast.CastMapStringInt(val); ok {
			return mval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return nil
}

// GetMapStringFloat returns the value of the path as map[string]float64, or nil if value of path is not a map or
// if any map keys are not castable to string or any values not castable to float64
func (gp *GPath) GetMapStringFloat(path string, fallback ...map[string]float64) map[string]float64 {
	if val, has := gp.get(path); has {
		if mval, ok := cast.CastMapStringFloat(val); ok {
			return mval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return nil
}

// GetMapStringBool returns the value of the path as map[string]bool, or nil if value of path is not a map or
// if any map keys are not castable to string or any values not castable to bool
func (gp *GPath) GetMapStringBool(path string, fallback ...map[string]bool) map[string]bool {
	if val, has := gp.get(path); has {
		if mval, ok := cast.CastMapStringBool(val); ok {
			return mval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return nil
}
