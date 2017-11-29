package gpath

import "github.com/ukautz/cast"

// IsIsBool returns bool whether path exists AND can be cast to bool
func (gp *GPath) IsBool(path string) bool {
	if val, has := gp.get(path); has {
		_, ok := cast.CastBool(val)
		return ok
	}
	return false
}

// GetBool returns the value of the path as bool, if it is a bool or can be casted into a bool
func (gp *GPath) GetBool(path string, fallback ...bool) bool {
	if val, has := gp.get(path); has {
		if bval, ok := cast.CastBool(val); ok {
			return bval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return false
}

// GetBools returns the value of the path as slice of bool, if it is a slice of bool or is a slice and each
// member can be casted into bool. Otherwise nil is returned.
func (gp *GPath) GetBools(path string, convertSingle ...bool) []bool {
	if val, has := gp.get(path); has {
		if res := cast.CastBools(val); res != nil {
			return res
		} else if len(convertSingle) > 0 && convertSingle[0] {
			if bval, ok := cast.CastBool(val); ok {
				return []bool{bval}
			}
		}
	}
	return nil
}
