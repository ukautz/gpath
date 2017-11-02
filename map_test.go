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
			from: map[interface{}]interface{}{"foo": "bar", "baz": 123, "zoing": []int{1, 3, 5}},
			key:  "foo",
			to:   &res{"bar", true},
		},
		{
			from: map[interface{}]interface{}{"foo": "bar", 234: 123, "zoing": []int{1, 3, 5}},
			key:  234,
			to:   &res{123, true},
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

func TestMapKeySet(t *testing.T) {
	ms := map[string]string{}
	//mi := map[string]interface{}
	assert.True(t, MapKeySet(ms, "foo", "bar"), "Set right key & right value")
	assert.Contains(t, ms, "foo", "Key found in map")
	assert.Equal(t, "bar", ms["foo"], "Value of key matches")

	assert.False(t, MapKeySet(ms, "bar", 123), "Set right key & wrong value fails")
	assert.NotContains(t, ms, "bar", "Unsettable key not found in map")

	assert.True(t, MapKeySet(ms, "bar", 123, true), "Set right key & wrong but castable value works")
	assert.Contains(t, ms, "bar", "Castable key found in map")
	assert.Equal(t, "123", ms["bar"], "Value of key set")

	assert.True(t, MapKeySet(ms, "bar", 123.23, true), "Set right key & wrong but castable value works")
	assert.Equal(t, "123.23", ms["bar"], "Value of key changed")

	assert.False(t, MapKeySet(ms, "bar", []float64{123.23}, true), "Set right key & wrong and uncastable value fails")
	assert.Equal(t, "123.23", ms["bar"], "Value of key unchanged")

	mi := map[interface{}]interface{}{}
	assert.True(t, MapKeySet(mi, "bar1", "foo"), "Set any key to interface should work (string)")
	assert.Contains(t, mi, "bar1", "Key created")

	assert.True(t, MapKeySet(mi, 123, "foo"), "Set any key to interface should work (int)")
	assert.Contains(t, mi, 123, "Key created")

	assert.True(t, MapKeySet(mi, 123.5, "foo"), "Set any key to interface should work (float)")
	assert.Contains(t, mi, 123.5, "Key created")

	m8 := map[int8]interface{}{}
	assert.True(t, MapKeySet(m8, "123", "foo", true), "Cast key works 1")
	assert.Contains(t, m8, int8(123), "Key created")

	assert.True(t, MapKeySet(m8, 33.33333, "foo", true), "Cast key works 2")
	assert.Contains(t, m8, int8(33), "Key created")

	assert.False(t, MapKeySet(m8, []int8{123}, "foo", true), "Uncastable key fails")
	assert.False(t, MapKeySet(m8, int16(123), "foo"), "Using wrong key fails")

	assert.True(t, MapKeySet(mi, 123.5, "foo"), "Set any key to interface should work (float)")
	assert.Contains(t, mi, 123.5, "Key created")

	assert.False(t, MapKeySet("foo", "bar", 123.23, true), "Setting on non map fails")
}
