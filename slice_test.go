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
