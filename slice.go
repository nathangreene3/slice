package slice

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

// Contains ...
func Contains(a []int, v int) bool {
	return true // Should depend on search or whatever it gets called
}

// Search or find or index?
func Search(a []int, v int) int {
	return 0
}

// RemoveAll of a value from a slice.
func RemoveAll(a []int, v int) []int {
	if len(a) == 0 {
		return a
	}

	// Iterate through a for i in [0,n), for n > 0.
	// If a[i] = v, then find smallest j in [i,n) such that a[j] != v.
	// If j = n, then a[i:n] = [v, v, ..., v], so return a[:i].
	// Otherwise, set a[i] as a[j], which isn't v.
	// Increment i and j.
	// Repeat while i and j are in [i,n).
	// When finished, i will be one larger than it should be, so return a[:i-1].
	var i, j int
	for ; i < len(a) && j < len(a); i, j = i+1, j+1 {
		if a[i] == v {
			for ; j < len(a) && a[j] == v; j++ {
			}

			if j == len(a) {
				return a[:i]
			}

			a[i] = a[j]
		}
	}

	return a[:i-1]
}
