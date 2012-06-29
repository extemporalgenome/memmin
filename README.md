# memmin

A trivial demonstration of (some of) Go's reflect capabilities.

This package provides a function, `Slice`, which takes an arbitrary slice type
and returns a value of that same type (boxed in an interface{}) that may or may
not have the property `len(slice) == cap(slice)`, depending on the provided
`abs` and `rel` parameters, which are an int and a float32, respectively.

`abs` is a threshold, as the integral capacity for extra elements past the
length, that the output slice should not exceed. Ex: an `abs` of 5 means that
`cap - len` should be less than or equal to 5.

`rel` is a threshold, as a ratio of 'extra' capacity, that the output slice
should not exceed. Ex: 0.25 means that the capacity should be no more than 25%
greater than the length.

If either of these thresholds is exceeded, the returned slice will point to a
freshly allocated copy that has no extra capacity. As a special case, if the
input slice is zero-length, a nil slice will always be returned. If neither
threshold is exceeded, the original slice will be returned.

`SlicePtr` is a convenience function; unlike `Slice`, which, as with `append`,
takes and returns a slice, this takes a pointer to a slice, which it modifies
in place, saving the need for unboxing.

## Usage notes

Should you use this? Probably not. While this does work, it is currently
intended to serve as a Go reflection demonstration. In the cases where this
behavior is useful (a very large slice of long-term data created via `append`),
it is better to spend the one or two lines of code to do this using the builtin
functions--you'll almost certainly know whether or not it's worth doing
(application dependent), and the builtins will be much faster.
