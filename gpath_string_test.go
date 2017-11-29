package gpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGPath_GetString(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path     string
		expect   string
		fallback bool
	}{
		{"string", "bar", false},
		{"strings", "", true},
		{"int", "123", false},
		{"ints", "", true},
		{"float", "12.5", false},
		{"floats", "", true},
		{"mixed-ok", "", true},
		{"mixed-ok.0", "1.5", false},
		{"mixed-ok.1", "2", false},
		{"mixed-ok.2", "3.5", false},
		{"mixed-nok", "", true},
		{"mixed-nok.0", "aaa", false},
		{"mixed-nok.1", "2", false},
		{"mixed-nok.2", "3.5", false},
		{"complex", "", true},
		{"complex.inner", "", true},
		{"complex.inner.0", "str", false},
		{"complex.inner.1", "123", false},
		{"complex.inner.2", "12.5", false},
		{"other", "", true},
		{"mixed-ok.3", "", true},
		{"mixed-nok.-1", "", true},
		{"mixed-nok.a", "", true},
		{"complex.other", "", true},
		{"complex.inner.4", "", true},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		val := gp.GetString(expect.path)
		valf := gp.GetString(expect.path, "__FALLBACK__")
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %###v", expect.path, expect.expect)
		if expect.fallback {
			assert.Equal(t, "__FALLBACK__", valf, "Path %s should fallback to 9999.5", expect.path)
		} else {
			assert.NotEqual(t, "__FALLBACK__", valf, "Path %s should NOT fallback to \"__FALLBACK__\"", expect.path)
		}
	}
}

func TestGPath_IsString(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path   string
		expect bool
	}{
		{"string", true},
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
		{"mixed-nok.0", true},
		{"mixed-nok.1", true},
		{"mixed-nok.2", true},
		{"complex", false},
		{"complex.inner", false},
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
		//fmt.Printf("E: %###v\n", expect)
		val := gp.IsString(expect.path)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect, val, "Path %s should be %v", expect.path, expect.expect)
	}
}

func TestGPath_GetStrings(t *testing.T) {
	gp := _newGPath()
	expects := []struct {
		path          string
		expect_strict []string
		expect_loose  []string
	}{
		{"string", nil, []string{"bar"}},
		{"strings", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"int", nil, []string{"123"}},
		{"ints", []string{"3", "4", "5"}, []string{"3", "4", "5"}},
		{"float", nil, []string{"12.5"}},
		{"floats", []string{"3.5", "4.5", "5.5"}, []string{"3.5", "4.5", "5.5"}},
		{"mixed-ok", []string{"1.5", "2", "3.5"}, []string{"1.5", "2", "3.5"}},
		{"mixed-ok.0", nil, []string{"1.5"}},
		{"mixed-ok.1", nil, []string{"2"}},
		{"mixed-ok.2", nil, []string{"3.5"}},
		{"mixed-nok", []string{"aaa", "2", "3.5"}, []string{"aaa", "2", "3.5"}},
		{"mixed-nok.0", nil, []string{"aaa"}},
		{"mixed-nok.1", nil, []string{"2"}},
		{"mixed-nok.2", nil, []string{"3.5"}},
		{"complex", nil, nil},
		{"complex.inner", []string{"str", "123", "12.5"}, []string{"str", "123", "12.5"}},
		{"complex.inner.0", nil, []string{"str"}},
		{"complex.inner.1", nil, []string{"123"}},
		{"complex.inner.2", nil, []string{"12.5"}},
		{"other", nil, nil},
		{"mixed-ok.3", nil, nil},
		{"mixed-nok.-1", nil, nil},
		{"mixed-nok.a", nil, nil},
		{"complex.other", nil, nil},
		{"complex.inner.4", nil, nil},
	}
	for _, expect := range expects {
		//fmt.Printf("E: %###v\n", expect)
		valStrict := gp.GetStrings(expect.path)
		valLoose := gp.GetStrings(expect.path, true)
		//fmt.Printf("V: %###v\n", val)
		assert.Equal(t, expect.expect_strict, valStrict, "Strict value from path %s should be %###v", expect.path, expect.expect_strict)
		assert.Equal(t, expect.expect_loose, valLoose, "Loose value from from path %s should be %###v", expect.path, expect.expect_loose)
	}
}
