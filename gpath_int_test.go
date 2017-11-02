package gpath

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGPath_IsInt(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path     string
		expect   bool
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
		val := gp.IsInt(expect.path)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_GetInt(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path     string
		expect   int64
		fallback bool
	}{
		{"string", 0, true},
		{"strings", 0, true},
		{"int", 123, false},
		{"ints", 0, true},
		{"float", 12, false},
		{"floats", 0, true},
		{"mixed-ok", 0, true},
		{"mixed-ok.0", 1, false},
		{"mixed-ok.1", 2, false},
		{"mixed-ok.2", 3, false},
		{"mixed-nok", 0, true},
		{"mixed-nok.0", 0, true},
		{"mixed-nok.1", 2, false},
		{"mixed-nok.2", 3, false},
		{"complex", 0, true},
		{"complex.inner", 0, true},
		{"complex.inner.0", 0, true},
		{"complex.inner.1", 123, false},
		{"complex.inner.2", 12, false},
		{"other", 0, true},
		{"mixed-ok.3", 0, true},
		{"mixed-nok.-1", 0, true},
		{"mixed-nok.a", 0, true},
		{"complex.other", 0, true},
		{"complex.inner.4", 0, true},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		val := gp.GetInt(expect.path)
		valf := gp.GetInt(expect.path, 9999)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %###v", expect.path, expect.expect)
		if expect.fallback {
			assert.Equal(t, int64(9999), valf, "Path %s should fallback to 9999.5", expect.path)
		} else {
			assert.NotEqual(t, int64(9999), valf, "Path %s should NOT fallback to 9999", expect.path)
		}
	}
}

func TestGPath_GetInts(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path          string
		expect_strict []int64
		expect_loose  []int64
	}{
		{"string", nil, nil},
		{"strings", nil, nil},
		{"int", nil, []int64{123}},
		{"ints", []int64{3, 4, 5}, []int64{3, 4, 5}},
		{"float", nil, []int64{12}},
		{"floats", []int64{3, 4, 5}, []int64{3, 4, 5}},
		{"mixed-ok", []int64{1, 2, 3}, []int64{1, 2, 3}},
		{"mixed-ok.0", nil, []int64{1}},
		{"mixed-ok.1", nil, []int64{2}},
		{"mixed-ok.2", nil, []int64{3}},
		{"mixed-nok", nil, nil},
		{"mixed-nok.0", nil, nil},
		{"mixed-nok.1", nil, []int64{2}},
		{"mixed-nok.2", nil, []int64{3}},
		{"complex", nil, nil},
		{"complex.inner", nil, nil},
		{"complex.inner.0", nil, nil},
		{"complex.inner.1", nil, []int64{123}},
		{"complex.inner.2", nil, []int64{12}},
		{"other", nil, nil},
		{"mixed-ok.3", nil, nil},
		{"mixed-nok.-1", nil, nil},
		{"mixed-nok.a", nil, nil},
		{"complex.other", nil, nil},
		{"complex.inner.4", nil, nil},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		valStrict := gp.GetInts(expect.path)
		valLoose := gp.GetInts(expect.path, true)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect_strict, valStrict, "Strict value from path %s should be %###v", expect.path, expect.expect_strict)
		assert.Equal(t, expect.expect_loose, valLoose, "Loose value from from path %s should be %###v", expect.path, expect.expect_loose)
	}
}