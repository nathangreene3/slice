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

func compare(a, b interface{}) int {
	switch at := a.(type) {
	case []bool:
		var (
			bt = b.([]bool)
			n  = math.MinInt(len(at), len(bt))
		)

		for i := 0; i < n; i++ {
			if at[i] {
				if !bt[i] {
					return 1
				}
			} else if bt[i] {
				return -1
			}
		}

		switch {
		case len(at) < len(bt):
			return -1
		case len(bt) < len(at):
			return 1
		default:
			return 0
		}
	case []float64:
		var (
			bt = b.([]float64)
			n  = math.MinInt(len(at), len(bt))
		)

		for i := 0; i < n; i++ {
			switch {
			case at[i] < bt[i]:
				return -1
			case at[i] < at[i]:
				return 1
			}
		}

		switch {
		case len(at) < len(bt):
			return -1
		case len(bt) < len(at):
			return 1
		default:
			return 0
		}
	case []int:
		var (
			bt = b.([]int)
			n  = math.MinInt(len(at), len(bt))
		)

		for i := 0; i < n; i++ {
			switch {
			case at[i] < bt[i]:
				return -1
			case at[i] < at[i]:
				return 1
			}
		}

		switch {
		case len(at) < len(bt):
			return -1
		case len(bt) < len(at):
			return 1
		default:
			return 0
		}

	case []int8:
		return 0
	case []int16:
		return 0
	case []int32:
		return 0
	case []int64:
		return 0
	case []string:
		return 0
	case []uint8:
		return 0
	case []uint16:
		return 0
	case []uint32:
		return 0
	case []uint64:
		return 0
	case []float32:
		return 0
	default:
		panic("unknown type")
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
	var (
		b = make([]int, 0, len(a))
		j int
	)

	for i := 0; i < len(a); i++ {
		for j = 0; j < len(vs) && a[i] != vs[j]; j++ {
		}

		if j == len(vs) {
			b = append(b, a[i])
		}
	}

	return b
}

// Resize a slice.
func Resize(a []int, n, c int) []int {
	b := make([]int, n, c)
	if len(a) < n {
		n = len(a)
	}

	copy(b, a[:n])
	return b
}

// Search for a value and return the smallest index
// referencing it. If not found, len(a) will be returned.
func Search(a []int, n int) int {
	var i int
	for ; i < len(a) && a[i] != n; i++ {
	}

	// TODO: Return -1 instead?
	return i
}

// SubSlice returns a copy of a[i:j].
func SubSlice(a []int, i, j int) []int {
	b := make([]int, j-i)
	copy(b, a[i:j])
	return b
}

// SubSliceSearch returns the index a subslice is found within a slice. If not
// found, len(a) is returned.
func SubSliceSearch(a, sub []int) int {
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

// Swap two values.
func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Unique returns a slice of the unique values in a given slice.
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

// unique returns a slice of the unique values in a given slice.
func unique(a []int) []int {
	// TODO: Compare this to Unique.
	b := make([]int, 0, cap(a))
	if 0 < len(a) {
		m := make(map[int]struct{})
		for i := 0; i < len(a); i++ {
			if _, ok := m[a[i]]; !ok {
				b = append(b, a[i])
				m[a[i]] = struct{}{}
			}
		}
	}

	return b
}
