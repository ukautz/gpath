package gpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGPath_IsBool(t *testing.T) {
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
		val := gp.IsBool(expect.path)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_GetBool(t *testing.T) {
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
		val := gp.GetBool(expect.path)
		valf := gp.GetBool(expect.path, true)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %###v", expect.path, expect.expect)
		assert.Equal(t, true, valf, "Path %s with fallback should be %###v", expect.path, expect.expect)
	}
}

func TestGPath_GetBools(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path          string
		expect_strict []bool
		expect_loose  []bool
	}{
		{"string", nil, nil},
		{"strings", nil, nil},
		{"int", nil, []bool{true}},
		{"ints", []bool{true, true, true}, []bool{true, true, true}},
		{"float", nil, []bool{true}},
		{"floats", []bool{true, true, true}, []bool{true, true, true}},
		{"mixed-ok", []bool{true, true, true}, []bool{true, true, true}},
		{"mixed-ok.0", nil, []bool{true}},
		{"mixed-ok.1", nil, []bool{true}},
		{"mixed-ok.2", nil, []bool{true}},
		{"mixed-nok", nil, nil},
		{"mixed-nok.0", nil, nil},
		{"mixed-nok.1", nil, []bool{true}},
		{"mixed-nok.2", nil, []bool{true}},
		{"complex", nil, nil},
		{"complex.inner", nil, nil},
		{"complex.inner.0", nil, nil},
		{"complex.inner.1", nil, []bool{true}},
		{"complex.inner.2", nil, []bool{true}},
		{"other", nil, nil},
		{"mixed-ok.3", nil, nil},
		{"mixed-nok.-1", nil, nil},
		{"mixed-nok.a", nil, nil},
		{"complex.other", nil, nil},
		{"complex.inner.4", nil, nil},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		valStrict := gp.GetBools(expect.path)
		valLoose := gp.GetBools(expect.path, true)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect_strict, valStrict, "Strict value from path %s should be %###v", expect.path, expect.expect_strict)
		assert.Equal(t, expect.expect_loose, valLoose, "Loose value from from path %s should be %###v", expect.path, expect.expect_loose)
	}
}
