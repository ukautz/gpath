package gpath

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cache(t *testing.T) {
	from := map[string]interface{}{"foo": 123}
	c := newCache(from)
	assert.Equal(t, from, c.data, "Initial data forwarded")

	val, ok := c.get("foo")
	assert.True(t, ok, "Read of init value succeeds")
	assert.Equal(t, 123, val, "Init value returned")

	val, ok = c.get("bar")
	assert.False(t, ok, "Read of not existing value fails")
	assert.Nil(t, val, "Not existing value is nil")

	v := 3.14159265359
	r := c.set("bar", v)
	assert.Equal(t, v, r, "Set returns value")

	assert.NotNil(t, c.set("zoing.aa", map[string]interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb1", map[string]interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb1.cc1", map[string]interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb1.cc1.dd1", map[string]interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb1.cc2", map[string]interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb2", map[string]interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb3", []interface{}{}))
	assert.NotNil(t, c.set("zoing.aa.bb3.0", "1"))
	assert.NotNil(t, c.set("zoing.aa.bb3.1", 2.22))
	assert.NotNil(t, c.set("zoing.aa.bb3.2", 3))
	assert.NotNil(t, c.set("zoing", map[string]interface{}{}))

	mustKeys := func(must bool, name string, keys []string) {
		for _, key := range keys {
			_, ok := c.get(key)
			if must {
				assert.True(t, ok, fmt.Sprintf("Key %s must exist @%s", key, name))
			} else {
				assert.False(t, ok, fmt.Sprintf("Key %s must NOT exist", key))
			}
		}
	}
	mustKeys(true, "clear1", []string{
		"zoing",
		"zoing.aa.bb1",
		"zoing.aa.bb1.cc1",
		"zoing.aa.bb1.cc1.dd1",
		"zoing.aa.bb1.cc2",
	})
	assert.Equal(t, 3, c.clear("zoing.aa.bb1."))
	mustKeys(false, "clear1", []string{
		"zoing.aa.bb1.cc1",
		"zoing.aa.bb1.cc1.dd1",
		"zoing.aa.bb1.cc2",
	})
	mustKeys(true, "clear1.1", []string{
		"zoing",
		"zoing.aa.bb1",
	})

	mustKeys(true, "clear2", []string{
		"zoing",
		"zoing.aa",
		"zoing.aa.bb1",
		"zoing.aa.bb2",
		"zoing.aa.bb3",
		"zoing.aa.bb3.0",
		"zoing.aa.bb3.1",
		"zoing.aa.bb3.2",
	})
	assert.Equal(t, 7, c.clear("zoing."))
	mustKeys(false, "clear2", []string{
		"zoing.aa",
		"zoing.aa.bb2",
		"zoing.aa.bb2",
		"zoing.aa.bb3",
		"zoing.aa.bb3.0",
		"zoing.aa.bb3.1",
		"zoing.aa.bb3.2",
	})
	mustKeys(true, "clear2.1", []string{
		"zoing",
	})

}
