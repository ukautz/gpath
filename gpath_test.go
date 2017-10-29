package gpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var _testData = map[string]interface{}{
	"string":    "bar",
	"strings":   []string{"a", "b", "c"},
	"int":       123,
	"ints":      []int{3, 4, 5},
	"float":     12.5,
	"floats":    []float32{3.5, 4.5, 5.5},
	"mixed-ok":  []interface{}{"1.5", uint(2), float32(3.5)},
	"mixed-nok": []interface{}{"aaa", uint(2), float32(3.5)},
	"complex": map[string]interface{}{
		"inner": []interface{}{
			"str",
			123,
			12.5,
		},
	},
}

func _newGPath() *GPath {
	return New(_testData)
}

func TestGPath_Has(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect bool
	}{
		{"string", true},
		{"strings", true},
		{"int", true},
		{"ints", true},
		{"float", true},
		{"floats", true},
		{"mixed-ok", true},
		{"mixed-ok.0", true},
		{"mixed-ok.1", true},
		{"mixed-ok.2", true},
		{"mixed-nok", true},
		{"mixed-nok.0", true},
		{"mixed-nok.1", true},
		{"mixed-nok.2", true},
		{"complex", true},
		{"complex.inner", true},
		{"complex.inner.0", true},
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
		assert.Equal(t, expect.expect, gp.Has(expect.path), "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_IsSlice(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect bool
	}{
		{"string", false},
		{"strings", true},
		{"int", false},
		{"ints", true},
		{"float", false},
		{"floats", true},
		{"mixed-ok", true},
		{"mixed-ok.0", false},
		{"mixed-ok.1", false},
		{"mixed-ok.2", false},
		{"mixed-nok", true},
		{"mixed-nok.0", false},
		{"mixed-nok.1", false},
		{"mixed-nok.2", false},
		{"complex", false},
		{"complex.inner", true},
		{"complex.inner.0", false},
		{"complex.inner.1", false},
		{"complex.inner.2", false},
		{"other", false},
		{"mixed-ok.3", false},
		{"mixed-nok.-1", false},
		{"mixed-nok.a", false},
		{"complex.other", false},
		{"complex.inner.4", false},
	}
	for _, expect := range expects {
		assert.Equal(t, expect.expect, gp.IsSlice(expect.path), "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_IsMap(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect bool
	}{
		{"string", false},
		{"strings", false},
		{"int", false},
		{"ints", false},
		{"float", false},
		{"floats", false},
		{"mixed-ok", false},
		{"mixed-ok.0", false},
		{"mixed-ok.1", false},
		{"mixed-ok.2", false},
		{"mixed-nok", false},
		{"mixed-nok.0", false},
		{"mixed-nok.1", false},
		{"mixed-nok.2", false},
		{"complex", true},
		{"complex.inner", false},
		{"complex.inner.0", false},
		{"complex.inner.1", false},
		{"complex.inner.2", false},
		{"other", false},
		{"mixed-ok.3", false},
		{"mixed-nok.-1", false},
		{"mixed-nok.a", false},
		{"complex.other", false},
		{"complex.inner.4", false},
	}
	for _, expect := range expects {
		assert.Equal(t, expect.expect, gp.IsMap(expect.path), "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_Get(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect interface{}
	}{
		{"string", "bar"},
		{"strings", []string{"a", "b", "c"}},
		{"int", 123},
		{"ints", []int{3, 4, 5}},
		{"float", 12.5},
		{"floats", []float32{3.5, 4.5, 5.5}},
		{"mixed-ok", []interface{}{"1.5", uint(2), float32(3.5)}},
		{"mixed-ok.0", "1.5"},
		{"mixed-ok.1", uint(2)},
		{"mixed-ok.2", float32(3.5)},
		{"mixed-nok", []interface{}{"aaa", uint(2), float32(3.5)}},
		{"mixed-nok.0", "aaa"},
		{"mixed-nok.1", uint(2)},
		{"mixed-nok.2", float32(3.5)},
		{"complex", map[string]interface{}{"inner": []interface{}{"str", 123, 12.5}}},
		{"complex.inner", []interface{}{"str", 123, 12.5}},
		{"complex.inner.0", "str"},
		{"complex.inner.1", 123},
		{"complex.inner.2", 12.5},
		{"other", nil},
		{"mixed-ok.3", nil},
		{"mixed-nok.-1", nil},
		{"mixed-nok.a", nil},
		{"complex.other", nil},
		{"complex.inner.4", nil},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		val := gp.Get(expect.path)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %###v", expect.path, expect.expect)
	}
}
