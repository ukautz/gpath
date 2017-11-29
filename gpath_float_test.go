package gpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGPath_IsFloat(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect bool
	}{
		{"string", false},
		{"strings", false},
		{"int", true},
		{"ints", false},
		{"float", true},
		{"floats", false},
		{"mixed-ok", false},
		{"mixed-ok.0", true},
		{"mixed-ok.1", true},
		{"mixed-ok.2", true},
		{"mixed-nok", false},
		{"mixed-nok.0", false},
		{"mixed-nok.1", true},
		{"mixed-nok.2", true},
		{"complex", false},
		{"complex.inner", false},
		{"complex.inner.0", false},
		{"complex.inner.1", true},
		{"complex.inner.2", true},
		{"other", false},
		{"mixed-ok.3", false},
		{"mixed-nok.-1", false},
		{"mixed-nok.a", false},
		{"complex.other", false},
		{"complex.inner.4", false},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		val := gp.IsFloat(expect.path)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_GetFloat(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path     string
		expect   float64
		fallback bool
	}{
		{"string", 0, true},
		{"strings", 0, true},
		{"int", 123, false},
		{"ints", 0, true},
		{"float", 12.5, false},
		{"floats", 0, true},
		{"mixed-ok", 0, true},
		{"mixed-ok.0", 1.5, false},
		{"mixed-ok.1", 2, false},
		{"mixed-ok.2", 3.5, false},
		{"mixed-nok", 0, true},
		{"mixed-nok.0", 0, true},
		{"mixed-nok.1", 2, false},
		{"mixed-nok.2", 3.5, false},
		{"complex", 0, true},
		{"complex.inner", 0, true},
		{"complex.inner.0", 0, true},
		{"complex.inner.1", 123, false},
		{"complex.inner.2", 12.5, false},
		{"other", 0, true},
		{"mixed-ok.3", 0, true},
		{"mixed-nok.-1", 0, true},
		{"mixed-nok.a", 0, true},
		{"complex.other", 0, true},
		{"complex.inner.4", 0, true},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		val := gp.GetFloat(expect.path)
		valf := gp.GetFloat(expect.path, 9999.5)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %###v", expect.path, expect.expect)
		if expect.fallback {
			assert.Equal(t, float64(9999.5), valf, "Path %s should fallback to 9999.5", expect.path)
		} else {
			assert.NotEqual(t, float64(9999.5), valf, "Path %s should NOT fallback to 9999.5", expect.path)
		}
	}
}

func TestGPath_GetFloats(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path          string
		expect_strict []float64
		expect_loose  []float64
	}{
		{"string", nil, nil},
		{"strings", nil, nil},
		{"int", nil, []float64{123}},
		{"ints", []float64{3, 4, 5}, []float64{3, 4, 5}},
		{"float", nil, []float64{12.5}},
		{"floats", []float64{3.5, 4.5, 5.5}, []float64{3.5, 4.5, 5.5}},
		{"mixed-ok", []float64{1.5, 2, 3.5}, []float64{1.5, 2, 3.5}},
		{"mixed-ok.0", nil, []float64{1.5}},
		{"mixed-ok.1", nil, []float64{2}},
		{"mixed-ok.2", nil, []float64{3.5}},
		{"mixed-nok", nil, nil},
		{"mixed-nok.0", nil, nil},
		{"mixed-nok.1", nil, []float64{2}},
		{"mixed-nok.2", nil, []float64{3.5}},
		{"complex", nil, nil},
		{"complex.inner", nil, nil},
		{"complex.inner.0", nil, nil},
		{"complex.inner.1", nil, []float64{123}},
		{"complex.inner.2", nil, []float64{12.5}},
		{"other", nil, nil},
		{"mixed-ok.3", nil, nil},
		{"mixed-nok.-1", nil, nil},
		{"mixed-nok.a", nil, nil},
		{"complex.other", nil, nil},
		{"complex.inner.4", nil, nil},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		valStrict := gp.GetFloats(expect.path)
		valLoose := gp.GetFloats(expect.path, true)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect_strict, valStrict, "Strict value from path %s should be %###v", expect.path, expect.expect_strict)
		assert.Equal(t, expect.expect_loose, valLoose, "Loose value from from path %s should be %###v", expect.path, expect.expect_loose)
	}
}
