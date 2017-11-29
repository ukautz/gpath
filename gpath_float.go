package gpath

import "github.com/ukautz/cast"

// IsFloat returns bool whether path exists AND can be cast to float (eg int(123), string("123.234") (=int(123)) or float64(123.234) (=int(123)))
func (gp *GPath) IsFloat(path string) bool {
	if val, has := gp.get(path); has {
		_, ok := cast.CastFloat(val)
		return ok
	}
	return false
}

// GetFloat returns the value of the path as float64, if it is a float64 or can be casted into a float64
func (gp *GPath) GetFloat(path string, fallback ...float64) float64 {
	if val, has := gp.get(path); has {
		if fval, ok := cast.CastFloat(val); ok {
			return fval
		}
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return 0
}

// GetFloats returns the value of the path as slice of float64, if it is a slice of float64 or is a slice and each
// member can be casted into float64. Otherwise nil is returned.
func (gp *GPath) GetFloats(path string, convertSingle ...bool) []float64 {
	if val, has := gp.get(path); has {
		if res := cast.CastFloats(val); res != nil {
			return res
		} else if len(convertSingle) > 0 && convertSingle[0] {
			if fval, ok := cast.CastFloat(val); ok {
				return []float64{fval}
			}
		}
	}
	return nil
}
