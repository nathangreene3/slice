package slice

import "strconv"

// Filterer defines the conditions for retaining a value.
type Filterer func(s string) bool

// Folder defines the accumulation of a value into a generic value.
type Folder func(x interface{}, s string) interface{}

// Generator defines the ith value in a slice.
type Generator func(i int) string

// Mapper defines the mapping of a value to a new value.
type Mapper func(s string) string

// Reducer defines the mapping (reduction) of two values to one.
type Reducer func(s, t string) string

// Filter values from a slice into a new slice.
func Filter(a []string, f Filterer) []string {
	b := make([]string, 0, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			b = append(b, a[i])
		}
	}

	return b
}

// Fold values into a single value given a seed x.
func Fold(a []string, x interface{}, f Folder) interface{} {
	y := x
	for i := 0; i < len(a); i++ {
		y = f(y, a[i])
	}

	return y
}

// Generate a slice.
func Generate(f Generator, n int) []string {
	a := make([]string, 0, n)
	for ; 0 < n; n-- {
		a = append(a, f(len(a)))
	}

	return a
}

// Map a slice to a new slice.
func Map(a []string, f Mapper) []string {
	b := make([]string, 0, len(a))
	for len(b) < len(a) {
		b = append(b, f(a[len(b)]))
	}

	return b
}

// Reduce a slice to a single value.
func Reduce(a []string, f Reducer) string {
	var v string
	for i := 0; i < len(a); i++ {
		v = f(v, a[i])
	}

	return v
}

// Compare two slices.
func Compare(a, b []string) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}

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

// Contains determines if any values are found in a slice.
func Contains(a []string, vs ...string) bool {
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
func Copy(a []string) []string {
	b := make([]string, len(a))
	copy(b, a)
	return b
}

// Count the number of occurances of a given number in a slice.
func Count(a []string, s string) int {
	var c int
	for i := 0; i < len(a); i++ {
		if a[i] == s {
			c++
		}
	}

	return c
}

// Equal determines if two slices have the same values at each index.
func Equal(a, b []string) bool {
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
func Freq(a []string) map[string]int {
	m := make(map[string]int)
	for i := 0; i < len(a); i++ {
		m[a[i]]++
	}

	return m
}

// Join several slices into one.
func Join(as ...[]string) []string {
	var n int
	for i := 0; i < len(as); i++ {
		n += len(as[i])
	}

	a := make([]string, 0, n)
	for i := 0; i < len(as); i++ {
		a = append(a, as[i]...)
	}

	return a
}

// Remove several values from a slice.
func Remove(a []string, vs ...string) []string {
	var (
		b = make([]string, 0, len(a))
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
func Resize(a []string, n, c int) []string {
	if len(a) < n {
		n = len(a)
	}

	b := make([]string, n, c)
	copy(b, a[:n])
	return b
}

// Search for a value and return the smallest index
// referencing it. If not found, len(a) will be returned.
func Search(a []string, s string) int {
	var i int
	for ; i < len(a) && a[i] != s; i++ {
	}

	return i
}

// SubSlice returns a copy of a[i:j].
func SubSlice(a []string, i, j int) []string {
	b := make([]string, j-i)
	copy(b, a[i:j])
	return b
}

// SearchSubSlice returns the index a subslice is found within a slice. If not
// found, len(a) is returned.
func SearchSubSlice(a, sub []string) int {
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
func Swap(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ToBts converts a slice of strings to a slice of bytes.
func ToBts(a []string) [][]byte {
	b := make([][]byte, 0, len(a))
	for i := 0; i < len(a); i++ {
		b = append(b, []byte(a[i]))
	}

	return b
}

// ToFlts converts a slice of strings to a slice of 64-bit floats.
func ToFlts(a []string) []float64 {
	b := make([]float64, 0, len(a))
	for i := 0; i < len(a); i++ {
		f, err := strconv.ParseFloat(a[i], 64)
		if err != nil {
			panic(err.Error())
		}

		b = append(b, f)
	}

	return b
}

// ToInts converts a slice of strings to a slice of 64-bit integers.
func ToInts(a []string, base int) []int {
	b := make([]int, 0, len(a))
	for i := 0; i < len(a); i++ {
		n, err := strconv.ParseInt(a[i], base, 64)
		if err != nil {
			panic(err.Error())
		}

		b = append(b, int(n))
	}

	return b
}

// ToMap returns a map of index-to-value pairs from a slice.
func ToMap(a []string) map[int]string {
	m := make(map[int]string)
	for len(m) < len(a) {
		m[len(m)] = a[len(m)]
	}

	return m
}

// ToUInts converts a slice to a slice of unsigned integers.
func ToUInts(a []string, base int) []uint {
	b := make([]uint, 0, len(a))
	for i := 0; i < len(a); i++ {
		n, err := strconv.ParseUint(a[i], base, 64)
		if err != nil {
			panic(err.Error())
		}

		b = append(b, uint(n))
	}

	return b
}

// Unique returns a slice containing the unique values in several slices.
func Unique(as ...[]string) []string {
	var n int
	for i := 0; i < len(as); i++ {
		n += len(as[i])
	}

	b := make([]string, 0, n)
	if 0 < n {
		m := make(map[string]struct{})
		for i := 0; i < len(as); i++ {
			for j := 0; j < len(as[i]); j++ {
				if _, ok := m[as[i][j]]; !ok {
					b = append(b, as[i][j])
					m[as[i][j]] = struct{}{}
				}
			}
		}
	}

	return b
}
