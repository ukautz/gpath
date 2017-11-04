package gpath

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

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

func TestGPath_GetMap(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect interface{}
	}{
		{"string", nil},
		{"strings", nil},
		{"int", nil},
		{"ints", nil},
		{"float", nil},
		{"floats", nil},
		{"mixed-ok", nil},
		{"mixed-ok.0", nil},
		{"mixed-ok.1", nil},
		{"mixed-ok.2", nil},
		{"mixed-nok", nil},
		{"mixed-nok.0", nil},
		{"mixed-nok.1", nil},
		{"mixed-nok.2", nil},
		{"complex", map[interface{}]interface{}{
			"inner": []interface{}{
				"str",
				123,
				12.5,
			},
		}},
		{"complex.inner", nil},
		{"complex.inner.0", nil},
		{"complex.inner.1", nil},
		{"complex.inner.2", nil},
		{"other", nil},
		{"mixed-ok.3", nil},
		{"mixed-nok.-1", nil},
		{"mixed-nok.a", nil},
		{"complex.other", nil},
		{"complex.inner.4", nil},
	}
	for _, expect := range expects {
		res := gp.GetMap(expect.path)
		if expect.expect == nil {
			assert.Nil(t, res, "Path %s should be nil but is %###v", expect.path, res)
		} else {
			assert.Equal(t, expect.expect, res, "Path %s should be %###v", expect.path, expect.expect)
			ress := gp.GetMapString(expect.path)
			assert.Equal(t, map[string]interface {}{"inner":[]interface {}{"str", 123, 12.5}}, ress, "Path %s should be that", expect.path)
		}

		resss := gp.GetMapStringString(expect.path)
		assert.Nil(t, resss, "Path %s should be nil but is %###v", expect.path, resss)

		ressi := gp.GetMapStringInt(expect.path)
		assert.Nil(t, ressi, "Path %s should be nil but is %###v", expect.path, ressi)

		ressf := gp.GetMapStringFloat(expect.path)
		assert.Nil(t, ressf, "Path %s should be nil but is %###v", expect.path, ressf)
	}
}

func TestGPath_GetMapString(t *testing.T) {
	source := map[string]interface{}{
		"ok1": map[string]interface{}{
			"foo": "bar",
		},
		"ok2": map[string]interface{}{
			"foo": map[string]interface{}{},
		},
	}
	gp := New(source)
	assert.Equal(t, map[string]interface{}{
		"foo": "bar",
	}, gp.GetMapString("ok1"))
	assert.Equal(t, map[string]interface{}{
		"foo": map[string]interface{}{},
	}, gp.GetMapString("ok2"))
	assert.Nil(t, gp.GetMapString("ok3"))
}

func TestGPath_GetMapStringString(t *testing.T) {
	source := map[string]interface{}{
		"ok": map[string]interface{}{
			"foo": "bar",
		},
		"nok": map[string]interface{}{
			"foo": map[string]interface{}{},
		},
	}
	gp := New(source)
	assert.Equal(t, map[string]string{
		"foo": "bar",
	}, gp.GetMapStringString("ok"))
	assert.Nil(t, gp.GetMapStringString("nok"))
	assert.Nil(t, gp.GetMapStringString("bla"))
}

func TestGPath_GetMapStringInt(t *testing.T) {
	source := map[string]interface{}{
		"ok": map[string]interface{}{
			"foo": "123.234",
		},
		"nok": map[string]interface{}{
			"foo": "bla",
		},
	}
	gp := New(source)
	assert.Equal(t, map[string]int64{
		"foo": 123,
	}, gp.GetMapStringInt("ok"))
	assert.Nil(t, gp.GetMapStringInt("nok"))
	assert.Nil(t, gp.GetMapStringInt("bla"))
}

func TestGPath_GetMapStringFloat(t *testing.T) {
	source := map[string]interface{}{
		"ok": map[string]interface{}{
			"foo": "123.5",
		},
		"nok": map[string]interface{}{
			"foo": "bla",
		},
	}
	gp := New(source)
	assert.Equal(t, map[string]float64{
		"foo": 123.5,
	}, gp.GetMapStringFloat("ok"))
	assert.Nil(t, gp.GetMapStringFloat("nok"))
	assert.Nil(t, gp.GetMapStringFloat("bla"))
}
