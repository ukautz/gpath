package gpath

import (
	"github.com/stretchr/testify/assert"
	"reflect"
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

func TestGPath_Set(t *testing.T) {
	type mi_t map[interface{}]interface{}
	gp := New(mi_t{})
	assert.Equal(t, reflect.Map, reflect.ValueOf(mi_t{}).Kind(), "type renaming keeps kind")

	assert.Nil(t, gp.Set("foo", "bar"), "can set root level key")
	assert.Contains(t, gp.source.(mi_t), "foo", "foo key was created")
	assert.Equal(t, "bar", gp.source.(mi_t)["foo"], "foo key was created")

	assert.NotNil(t, gp.Set("bar.baz", 234), "cannot set deep structure with missing parents")

	assert.Nil(t, gp.Set("bar", mi_t{"baz": 234}), "cannot set deep structure with missing parents")
	assert.Contains(t, gp.source.(mi_t), "bar", "bar key was created")
	assert.Contains(t, gp.source.(mi_t)["bar"], "baz", "bar.baz key was created")
	assert.Equal(t, 234, gp.source.(mi_t)["bar"].(mi_t)["baz"], "foo key was created")

	assert.Nil(t, gp.Set("bar.zoing", 345), "can now set other child elements")
	assert.Contains(t, gp.source.(mi_t)["bar"], "zoing", "bar.zoing key was created")
	assert.Equal(t, 345, gp.source.(mi_t)["bar"].(mi_t)["zoing"], "foo key was created")

	gp = New(mi_t{"strs": []string{}})
	assert.Nil(t, gp.Set("strs.-1", "foo"), "can add slice element")
	assert.Contains(t, gp.source.(mi_t), "strs", "parent is still here")
	assert.Equal(t, gp.source.(mi_t)["strs"], []string{"foo"}, "slice value added")

	gp = New(mi_t{"iii": []interface{}{[]interface{}{[]interface{}{"foo"}}}})
	assert.Nil(t, gp.Set("iii.0.0.-1", "bar"), "can add slice within slices element")
	assert.Contains(t, gp.source.(mi_t), "iii", "parent is still here")
	assert.Equal(t, []interface{}{[]interface{}{[]interface{}{"foo", "bar"}}}, gp.source.(mi_t)["iii"], "foo was created")

	ss := []string{}
	gp = New(&ss)
	assert.Nil(t, gp.Set("-1", "bar"), "can add to indirect slice ref")
	assert.Equal(t, &[]string{"bar"}, gp.source.(*[]string), "added to slice ref")

	gp = New(ss)
	assert.NotNil(t, gp.Set("-1", "bar"), "can NOT add to direct slice")

	gp = New(&ss)
	assert.NotNil(t, gp.Set("key", "bar"), "can NOT set string key in slice ref")

	gp = New("string")
	assert.NotNil(t, gp.Set("key", "bar"), "can NOT set in scalar")
}
