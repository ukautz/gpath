package gpath

import (
	"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
)

func TestSliceIndex(t *testing.T) {
	type res struct {
		val   interface{}
		found bool
	}
	type expect struct {
		from interface{}
		idx  int
		to   *res
	}
	expects := []*expect{
		{
			from: []string{"foo", "bar", "baz"},
			idx:  -1,
			to:   &res{nil, false},
		},
		{
			from: []string{"foo", "bar", "baz"},
			idx:  0,
			to:   &res{"foo", true},
		},
		{
			from: []string{"foo", "bar", "baz"},
			idx:  1,
			to:   &res{"bar", true},
		},
		{
			from: []string{"foo", "bar", "baz"},
			idx:  2,
			to:   &res{"baz", true},
		},
		{
			from: []string{"foo", "bar", "baz"},
			idx:  3,
			to:   &res{nil, false},
		},
		{
			from: "not-slice",
			idx:  0,
			to:   &res{nil, false},
		},
		{
			from: 123,
			idx:  0,
			to:   &res{nil, false},
		},
		{
			from: map[string]interface{}{"foo": "bar"},
			idx:  0,
			to:   &res{nil, false},
		},
		{
			from: map[int]string{1: "bar"},
			idx:  1,
			to:   &res{nil, false},
		},
	}

	for _, e := range expects {
		res, found := SliceIndex(e.from, e.idx)
		assert.Equal(t, e.to.found, found, fmt.Sprintf("Expect found = %v (idx: %d, from %###v)", e.to.found, e.idx, e.from))
		if e.to.found {
			assert.Equal(t, e.to.val, res, fmt.Sprintf("Result is %v (idx: %d, from %###v)", e.to.val, e.idx, e.from))
		}
	}
}

func TestSliceIndexSet(t *testing.T) {
	ss := []string{}
	assert.True(t, SliceIndexSet(&ss, 0, "abc"), "Appending of right value should work")
	assert.Equal(t, []string{"abc"}, ss, "Appended string is element")

	assert.True(t, SliceIndexSet(&ss, 0, "cde"), "Replacing of right value should work")
	assert.Equal(t, []string{"cde"}, ss, "Replaced string is element")

	assert.True(t, SliceIndexSet(&ss, -1, "bla"), "Auto-append of right value should work")
	assert.Equal(t, []string{"cde", "bla"}, ss, "Auto-appended string is element")

	assert.False(t, SliceIndexSet(&ss, -1, 123), "Appending wrong type fails")
	assert.Equal(t, []string{"cde", "bla"}, ss, "Wrong type element not in")

	assert.True(t, SliceIndexSet(&ss, -1, 123, true), "Appending wrong type with casting works")
	assert.Equal(t, []string{"cde", "bla", "123"}, ss, "Wrong type with casting is casted element")

	assert.False(t, SliceIndexSet(ss, -1, "foo", true), "Appending right type on slice, not pointer to slice, fails")
	assert.Equal(t, []string{"cde", "bla", "123"}, ss, "Unchanged after trying non-pointer")

	assert.False(t, SliceIndexSet(&ss, -3, "foo"), "Setting to invalid negative index fails")
	assert.Equal(t, []string{"cde", "bla", "123"}, ss, "Unchanged after trying invalid negative index")

	assert.False(t, SliceIndexSet(&ss, 100, "foo"), "Setting to invalid positive index fails")
	assert.Equal(t, []string{"cde", "bla", "123"}, ss, "Unchanged after trying invalid positive index")

	assert.False(t, SliceIndexSet(&ss, -1, []map[string]interface{}{}, true), "Uncastable value cannot be appended")
	assert.Equal(t, []string{"cde", "bla", "123"}, ss, "Unchanged after trying top add uncastable value")


	ii := []interface{}{}
	assert.True(t, SliceIndexSet(&ii, -1, "foo"), "Interface slice should accept anything (string)")
	assert.Equal(t, []interface{}{"foo"}, ii, "String value added to interface slice")

	assert.True(t, SliceIndexSet(&ii, -1, 123), "Interface slice should accept anything (int)")
	assert.Equal(t, []interface{}{"foo", 123}, ii, "Int value added to interface slice")

	assert.True(t, SliceIndexSet(&ii, -1, 22.5), "Interface slice should accept anything (float)")
	assert.Equal(t, []interface{}{"foo", 123, 22.5}, ii, "Float value added to interface slice")

	assert.True(t, SliceIndexSet(&ii, -1, []interface{}{999}), "Interface slice should accept anything (interface slice)")
	assert.Equal(t, []interface{}{"foo", 123, 22.5, []interface{}{999}}, ii, "Interface slice value added to interface slice")


	assert.False(t, SliceIndexSet("bla", -1, "foo"), "Appending to not-slice (string) fails")
	assert.False(t, SliceIndexSet(map[string]interface{}{}, -1, "foo"), "Appending to not-slice (map) fails")
}