package gpath


import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapKey(t *testing.T) {
	type res struct {
		val   interface{}
		found bool
	}
	type expect struct {
		from interface{}
		key  interface{}
		to   *res
	}
	expects := []*expect{
		{
			from: map[string]interface{}{"foo": "bar", "baz": 123, "zoing": []int{1, 3, 5}},
			key:  "foo",
			to:   &res{"bar", true},
		},
		{
			from: map[string]interface{}{"foo": "bar", "baz": 123, "zoing": []int{1, 3, 5}},
			key:  "baz",
			to:   &res{123, true},
		},
		{
			from: map[string]interface{}{"foo": "bar", "baz": 123, "zoing": []int{1, 3, 5}},
			key:  "zoing",
			to:   &res{[]int{1, 3, 5}, true},
		},
		{
			from: map[string]interface{}{"foo": "bar", "baz": 123, "zoing": []int{1, 3, 5}},
			key:  "abc",
			to:   &res{nil, false},
		},
		{
			from: map[string]interface{}{"foo": "bar", "baz": 123, "zoing": []int{1, 3, 5}},
			key:  123,
			to:   &res{nil, false},
		},
		{
			from: []string{"foo", "bar", "baz"},
			key:  1,
			to:   &res{nil, false},
		},
		{
			from: "not-map",
			key:  0,
			to:   &res{nil, false},
		},
		{
			from: 123,
			key:  0,
			to:   &res{nil, false},
		},
	}

	for _, e := range expects {
		res, found := MapKey(e.from, e.key)
		assert.Equal(t, e.to.found, found, fmt.Sprintf("Expect found = %v (key: %v, from %###v)", e.to.found, e.key, e.from))
		if e.to.found {
			assert.Equal(t, e.to.val, res, fmt.Sprintf("Result is %v (key: %v, from %###v)", e.to.val, e.key, e.from))
		}
	}
}
