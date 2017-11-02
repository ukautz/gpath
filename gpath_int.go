package gpath

import (
	"github.com/ukautz/cast"
)

// IsInt returns bool whether path exists AND can be cast to int (eg int(123), string("123.234") (=int(123)) or float64(123.234) (=int(123)))
func (gp *GPath) IsInt(path string) bool {
	if val, has := gp.get(path); has {
		_, ok := cast.CastInt(val)
		return ok
	}
	return false
}

// GetInt returns the value of the path as int64, if it is a int64 or can be casted into a int64
func (gp *GPath) GetInt(path string, fallback ...int64) int64 {
	if val, has := gp.get(path); has {
		if ival, ok := cast.CastInt(val); ok {
			return ival
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return 0
}

// GetInts returns the value of the path as slice of int64, if it is a slice of int64 or is a slice and each
// member can be casted into int64. Otherwise nil is returned.
func (gp *GPath) GetInts(path string, convertSingle ...bool) []int64 {
	if val, has := gp.get(path); has {
		if res := cast.CastInts(val); res != nil {
			return res
		} else if len(convertSingle) > 0 && convertSingle[0] {
			if ival, ok := cast.CastInt(val); ok {
				return []int64{ival}
			}
		}
	}
	return nil
}
