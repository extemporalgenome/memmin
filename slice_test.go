// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memmin

import "testing"

func data() []int {
	return []int{1, 7, 4, 2, 0, 0, 0, 0}[:4]
}

var tests = []struct {
	abs int
	rel float32
	new bool
}{
	{0, 0, true},
	{0, -1, true},
	{-1, 0, true},
	{4, 1.0, false},
	{4, -1, false},
	{-1, 1.0, false},
	{3, -1, true},
	{-1, 0.99, true},
	{3, 0.99, true},
	{4, 0.99, true},
	{3, 1.0, true},
}

func TestSlice(t *testing.T) {
	src := data()
	for i, test := range tests {
		dst := Slice(src, test.abs, test.rel).([]int)
		if len(dst) == cap(dst) != test.new {
			t.Fatal("Failed on test %d: (%v)\n", i, test)
		}
	}
}

func TestSlicePtr(t *testing.T) {
	for i, test := range tests {
		d := data()
		SlicePtr(&d, test.abs, test.rel)
		if len(d) == cap(d) != test.new {
			t.Fatal("Failed on test %d: (%v)\n", i, test)
		}
	}
}
