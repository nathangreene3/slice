package slice

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
	for i := 0; i < len(a); i++ {
		b = append(b, f(a[i]))
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
	for i := 0; i < len(a); i++ {
		m[i] = a[i]
	}

	return m
}

// Contains determines if a value is contained within a slice.
func Contains(a []int, v int) bool {
	return Search(a, v) < len(a)
}

// Search for a value and return the smallest index
// referencing it. If not found, n will be returned.
func Search(a []int, v int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == v {
			return i
		}
	}

	// TODO: Should this return -1 instead?
	return len(a)
}

// RemoveAll of a value from a slice.
func RemoveAll(a []int, v int) []int {
	return Filter(a, func(ai int) bool { return ai != v })
}

// Resize a slice.
func Resize(a []int, n, c int) []int {
	b := make([]int, n, c)
	if n < len(a) {
		copy(b, a[:n])
	} else {
		copy(b, a)
	}

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
		if v, ok := m[ai]; !ok {
			m[ai] = v
			return true
		}

		return false
	}

	return Filter(a, f)
}
