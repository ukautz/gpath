[![Build Status](https://travis-ci.org/ukautz/gpath.svg?branch=master)](https://travis-ci.org/ukautz/gpath)
[![Coverage](http://gocover.io/_badge/github.com/ukautz/gpath?v=0.1.1)](http://gocover.io/github.com/ukautz/gpath)
[![GoDoc](https://godoc.org/github.com/ukautz/gpath?status.svg)](https://godoc.org/github.com/ukautz/gpath)

GPath
=====

GPath provides path based access to map or slice data structures in Go. Additionally type conversion helper methods are provided, helping to work with user input. Basically

* Use `users.0.id` (`<key|idx>[.<key|idx>[.<key|idx>[...]]]`) path notation to easily access complex Go data structures (eg configuration files, arbitrary JSON/YAML/... data, ..)
* Fast high level API (`Has(path)`, `GetString(path[, fallback])`, `IsSlice`, `GetFloats`, ..) with underlying cache

```go
package example

import "github.com/ukautz/gpath"

func example() {
	
	// have some data
	data := map[string]interface{}{
		"foo1": "bar",
		"foo2": 123,
		"foo3": 33.44,
		"baz": []interface{}{
			map[string]interface{}{
				"bla": "blub",
			},
			123.45,
			[]string{"a", "b", "c"},
			[]int{2, 3, 4},
			[]float32{1.1, 2.2, 3.3},
		},
	}
	
	// read it
	gp := gpath.New(data)
	
	// get primitive types
	s := gp.GetString("foo1") // string("bar")
	i := gp.GetInt("foo2") // int64(123)
	f := gp.GetFloat("foo3") // float64(33.44)
	
	// .. or convert from other primitive types
	s = gp.GetString("foo3") // string("33.44")
	i = gp.GetInt("foo3") // int64(33)
	f = gp.GetFloat("foo2") // float64(123)
	
	// .. as long as you can convert them
	s = gp.GetString("baz") // string("")
	i = gp.GetInt("foo1") // int64(0)
	f = gp.GetFloat("foo1") // float64(0)
	
	// .. within deep structures
	s = gp.GetString("baz.0.bla") // string("blub")
	f = gp.GetFloat("baz.1") // float64(123.45)
	
	// .. also with slices
	ss := gp.GetStrings("baz.2") // []string{"a", "b", "c"}
	is := gp.GetInts("baz.3") // []int64{2, 3, 4}
	fs := gp.GetFloats("baz.4") // []float64{1.1, 2.2, 3.3}
	
	// .. which support conversion as well
	ss = gp.GetStrings("baz.4") // []string{"1.1", "2.2", "3.3"}
	is = gp.GetInts("baz.4") // []int64{1, 2, 3}
	
	// .. as long as it's possible
	fs = gp.GetFloats("baz.2") // nil
}
```

Installation
------------

```bash
go get github.com/ukautz/gpath
```

This package supports version tags for [package management tools](https://github.com/golang/go/wiki/PackageManagementTools).