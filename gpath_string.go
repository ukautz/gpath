package gpath

import (
	"github.com/ukautz/cast"
)

// IsString returns bool whether path exists AND can be cast to string (eg actual string, int, or float)
func (gp *GPath) IsString(path string) bool {
	if val, has := gp.get(path); has {
		_, ok := cast.CastString(val)
		return ok
	}
	return false
}

// GetString returns the value of the path as string, if it is a string or can be casted into a string
func (gp *GPath) GetString(path string, fallback ...string) string {
	if val, has := gp.get(path); has {
		if sval, ok := cast.CastString(val); ok {
			return sval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return ""
}

// GetStrings returns the value of the path as slice of string, if it is a slice of string or is a slice and each
// member can be casted into string. Otherwise nil is returned.
func (gp *GPath) GetStrings(path string, convertSingle ...bool) []string {
	if val, has := gp.get(path); has {
		if res := cast.CastStrings(val); res != nil {
			return res
		} else if len(convertSingle) > 0 && convertSingle[0] {
			if sval, ok := cast.CastString(val); ok {
				return []string{sval}
			}
		}
	}
	return nil
}
