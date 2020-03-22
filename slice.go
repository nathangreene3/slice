package slice

import "github.com/nathangreene3/math"

// ------------------------------------------------------------
// DESIGN DECISION
// ------------------------------------------------------------
// If a function maps a type to the same type, the given type
// will not be mutated and the returned type will be a new
// value. If no type is returned, then the given type is
// potentially mutated.
// ------------------------------------------------------------

// Filterer defines the conditions for retaining a value.
type Filterer func(n int) bool

// Generator defines the ith value in a slice.
type Generator func(i int) int

// Mapper defines the mapping of a value to a new value.
type Mapper func(n int) int

// Reducer defines the mapping (reduction) of two values to one.
type Reducer func(m, n int) int

// Filter values from a slice into a new slice.
func Filter(a []int, f Filterer) []int {
	b := make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			b = append(b, a[i])
		}
	}

	return b
}

// Generate a slice.
func Generate(f Generator, n int) []int {
	a := make([]int, 0, n)
	for ; 0 < n; n-- {
		a = append(a, f(len(a)))
	}

	return a
}

// Map a slice to a new slice.
func Map(a []int, f Mapper) []int {
	b := make([]int, 0, len(a))
	for len(b) < len(a) {
		b = append(b, f(a[len(b)]))
	}

	return b
}

// Reduce a slice to a single value.
func Reduce(a []int, f Reducer) int {
	var v int
	for i := 0; i < len(a); i++ {
		v = f(v, a[i])
	}

	return v
}

// ToMap returns a map of index-to-value pairs from a slice.
func ToMap(a []int) map[int]int {
	m := make(map[int]int)
	for len(m) < len(a) {
		m[len(m)] = a[len(m)]
	}

	return m
}

// Compare ...
func Compare(a, b []int) int {
	n := math.MinInt(len(a), len(b))
	for i := 0; i < n; i++ {
		switch {
		case a[i] < b[i]:
			return -1
		case b[i] < a[i]:
			return 1
		}
	}

	switch {
	case len(a) < len(b):
		return -1
	case len(b) < len(a):
		return 1
	default:
		return 0
	}
}

// Contains ...
func Contains(a []int, vs ...int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(vs); j++ {
			if a[i] == vs[j] {
				return true
			}
		}
	}

	return false
}

// ContainsSubSlice returns the index a subslice is found within a slice. If not
// found, len(a) is returned.
func ContainsSubSlice(a, sub []int) int {
	if 0 < len(a) && 0 < len(sub) {
		for i := 0; i+len(sub) < len(a); i++ {
			if a[i] == sub[0] {
				for j := 1; j < len(sub); j++ {
					if a[i+j] != sub[j] {
						return len(a)
					}
				}

				return i
			}
		}
	}

	return len(a)
}

// Copy a slice.
func Copy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

// Count the number of occurances of a given number in a slice.
func Count(a []int, n int) int {
	var c int
	for i := 0; i < len(a); i++ {
		if a[i] == n {
			c++
		}
	}

	return c
}

// Equal ...
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

// Freq returns the mapping of each value in a slice to its
// frequency of occurance in the slice.
func Freq(a []int) map[int]int {
	m := make(map[int]int)
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}

	return m
}

// Join several slices into one.
func Join(as ...[]int) []int {
	var n int
	for i := 0; i < len(as); i++ {
		n += len(as[i])
	}

	a := make([]int, n)
	n = 0
	for i := 0; i < len(as); i++ {
		n += copy(a[n:n+len(as[i])], as[i])
	}

	return a
}

// Join2 ...
func Join2(as ...[]int) []int {
	var n int
	for i := 0; i < len(as); i++ {
		n += len(as[i])
	}

	a := make([]int, 0, n)
	for i := 0; i < len(as); i++ {
		a = append(a, as[i]...)
	}

	return a
}

// Remove several values from a slice.
func Remove(a []int, vs ...int) []int {
	f := func(ai int) bool {
		for j := 0; j < len(vs); j++ {
			if ai == vs[j] {
				return false
			}
		}

		return true
	}

	return Filter(a, f)
}

// Resize a slice.
func Resize(a []int, n, c int) []int {
	b := make([]int, n, c)
	copy(b, a[:math.MinInt(n, len(a))])
	return b
}

// Search for a value and return the smallest index
// referencing it. If not found, n will be returned.
func Search(a []int, n int) int {
	var i int
	for ; i < len(a) && a[i] != n; i++ {
	}

	// TODO: Should this return -1 instead?
	return i
}

// SubSlice returns a copy of a[i:j].
func SubSlice(a []int, i, j int) []int {
	b := make([]int, j-i)
	copy(b, a[i:j])
	return b
}

// Swap two values.
func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Unique filters a for unique values. This function is stable.
func Unique(a []int) []int {
	m := make(map[int]struct{})
	f := func(ai int) bool {
		_, ok := m[ai]
		if !ok {
			m[ai] = struct{}{}
		}

		return !ok
	}

	return Filter(a, f)
}
