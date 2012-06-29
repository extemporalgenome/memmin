// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memmin

import (
	"reflect"
)

type SliceType interface{}
type SlicePtrType interface{}

func Slice(slice SliceType, abs int, rel float32) SliceType {
	v, _ := handleslice(reflect.ValueOf(slice), abs, rel)
	return v.Interface()
}

func SlicePtr(ptr SlicePtrType, abs int, rel float32) (changed bool) {
	p := reflect.ValueOf(ptr)
	v0 := p.Elem()
	if v1, modified := handleslice(v0, abs, rel); modified {
		v0.Set(v1)
		changed = true
	}
	return changed
}

func handleslice(v reflect.Value, abs int, rel float32) (reflect.Value, bool) {
	len := v.Len()
	cap := v.Cap()
	changed := true
	switch {
	case len == 0:
		v = reflect.Zero(v.Type())
	case abs >= 0 && cap-len > abs:
		fallthrough
	case rel >= 0 && float32(cap-len)/float32(len) > rel:
		v = makeslice(v, len, len)
	default:
		changed = false
	}
	return v, changed
}

func makeslice(v reflect.Value, len, cap int) reflect.Value {
	dst := reflect.MakeSlice(v.Type(), len, cap)
	reflect.Copy(dst, v)
	return dst
}
