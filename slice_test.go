package slice

import "testing"

// Factor a positive integer. This is the naive solution.
func Factor(n int) []int {
	return Filter(
		Generate(func(i int) int { return i + 1 }, n),
		func(ai int) bool { return n%ai == 0 },
	)
}

// Prime determines if a positive integer is prime.
func Prime(n int) bool {
	return len(Factor(n)) == 2
}

// Primes returns a slice of primes less than or equal to n.
func Primes(n int) []int {
	return Filter(
		Generate(func(i int) int { return i + 1 }, n),
		func(ai int) bool { return Prime(ai) },
	)
}

func TestPrimes(t *testing.T) {
	// t.Fatal(Primes(100))
}

func TestCompare(t *testing.T) {
	tests := []struct {
		a, b []int
		exp  int
	}{
		// Ensure test cases here are also found in TestEqual

		// Permutations of slices
		{
			a:   []int{},
			b:   []int{},
			exp: 0,
		},
		{
			a:   []int{1},
			b:   []int{1},
			exp: 0,
		},
		{
			a:   []int{1, 2},
			b:   []int{1, 2},
			exp: 0,
		},
		{
			a:   []int{1, 2},
			b:   []int{2, 1},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1, 2, 3},
			exp: 0,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1, 3, 2},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{3, 1, 2},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{2, 1, 3},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{2, 3, 1},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{3, 2, 1},
			exp: -1,
		},

		// Equal lengths, unequal values
		{
			a:   []int{1},
			b:   []int{2},
			exp: -1,
		},
		{
			a:   []int{1, 2},
			b:   []int{2, 3},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{2, 3, 4},
			exp: -1,
		},

		// Unequal lengths
		{
			a:   []int{},
			b:   []int{1},
			exp: -1,
		},
		{
			a:   []int{1},
			b:   []int{},
			exp: 1,
		},
		{
			a:   []int{},
			b:   []int{1, 2},
			exp: -1,
		},
		{
			a:   []int{1, 2},
			b:   []int{},
			exp: 1,
		},
		{
			a:   []int{1},
			b:   []int{1, 2},
			exp: -1,
		},
		{
			a:   []int{1, 2},
			b:   []int{1},
			exp: 1,
		},
		{
			a:   []int{},
			b:   []int{1, 2, 3},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{},
			exp: 1,
		},
		{
			a:   []int{1},
			b:   []int{1, 2, 3},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1},
			exp: 1,
		},
		{
			a:   []int{1, 2},
			b:   []int{1, 2, 3},
			exp: -1,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1, 2},
			exp: 1,
		},
	}

	for _, test := range tests {
		if rec := Compare(test.a, test.b); test.exp != rec {
			t.Fatalf("\n   given: (%d, %d)\nexpected: %d\nreceived: %d\n", test.a, test.b, test.exp, rec)
		}
	}
}

func TestContains(t *testing.T) {
	// TODO: Convert v int to vs []int
	tests := []struct {
		a   []int
		v   int
		exp bool
	}{
		{
			a:   []int{},
			v:   1,
			exp: false,
		},
		{
			a:   []int{1},
			v:   1,
			exp: true,
		},
		{
			a:   []int{1},
			v:   2,
			exp: false,
		},
		{
			a:   []int{1, 2},
			v:   1,
			exp: true,
		},
		{
			a:   []int{2, 1},
			v:   1,
			exp: true,
		},
		{
			a:   []int{1, 2},
			v:   3,
			exp: false,
		},
		{
			a:   []int{2, 1},
			v:   3,
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			v:   1,
			exp: true,
		},
		{
			a:   []int{1, 3, 2},
			v:   1,
			exp: true,
		},
		{
			a:   []int{3, 1, 2},
			v:   1,
			exp: true,
		},
		{
			a:   []int{2, 1, 3},
			v:   1,
			exp: true,
		},
		{
			a:   []int{2, 3, 1},
			v:   1,
			exp: true,
		},
		{
			a:   []int{3, 2, 1},
			v:   1,
			exp: true,
		},
		{
			a:   []int{1, 2, 3},
			v:   4,
			exp: false,
		},
		{
			a:   []int{1, 3, 2},
			v:   4,
			exp: false,
		},
		{
			a:   []int{3, 1, 2},
			v:   4,
			exp: false,
		},
		{
			a:   []int{2, 1, 3},
			v:   4,
			exp: false,
		},
		{
			a:   []int{2, 3, 1},
			v:   4,
			exp: false,
		},
		{
			a:   []int{3, 2, 1},
			v:   4,
			exp: false,
		},
	}

	for _, test := range tests {
		if rec := Contains(test.a, test.v); test.exp != rec {
			t.Fatalf("\n   given: %d\nexpected: %t\nreceived: %t\n", test.a, test.exp, rec)
		}
	}
}

func TestContainsSubSlice(t *testing.T) {
	// TODO
}

func TestCount(t *testing.T) {
	tests := []struct {
		a      []int
		n, exp int
	}{
		{
			a:   []int{},
			n:   1,
			exp: 0,
		},
		{
			a:   []int{1},
			n:   2,
			exp: 0,
		},
		{
			a:   []int{1, 2},
			n:   3,
			exp: 0,
		},
		{
			a:   []int{1, 2, 3},
			n:   4,
			exp: 0,
		},
		{
			a:   []int{1},
			n:   1,
			exp: 1,
		},
		{
			a:   []int{1, 2},
			n:   1,
			exp: 1,
		},
		{
			a:   []int{2, 1},
			n:   1,
			exp: 1,
		},
		{
			a:   []int{1, 1, 2},
			n:   1,
			exp: 2,
		},
		{
			a:   []int{1, 2, 1},
			n:   1,
			exp: 2,
		},
		{
			a:   []int{2, 1, 1},
			n:   1,
			exp: 2,
		},
	}

	for _, test := range tests {
		if rec := Count(test.a, test.n); test.exp != rec {
			t.Fatalf("\n   given: (%d, %d)\nexpected: %d\nreceived: %d\n", test.a, test.n, test.exp, rec)
		}
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		a, b []int
		exp  bool
	}{
		// Ensure test cases here are also found in TestCompare

		// Permutations of slices
		{
			a:   []int{},
			b:   []int{},
			exp: true,
		},
		{
			a:   []int{1},
			b:   []int{1},
			exp: true,
		},
		{
			a:   []int{1, 2},
			b:   []int{1, 2},
			exp: true,
		},
		{
			a:   []int{1, 2},
			b:   []int{2, 1},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1, 2, 3},
			exp: true,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1, 3, 2},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{3, 1, 2},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{2, 1, 3},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{2, 3, 1},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{3, 2, 1},
			exp: false,
		},

		// Equal lengths, unequal values
		{
			a:   []int{1},
			b:   []int{2},
			exp: false,
		},
		{
			a:   []int{1, 2},
			b:   []int{2, 3},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{2, 3, 4},
			exp: false,
		},

		// Unequal lengths
		{
			a:   []int{},
			b:   []int{1},
			exp: false,
		},
		{
			a:   []int{1},
			b:   []int{},
			exp: false,
		},
		{
			a:   []int{},
			b:   []int{1, 2},
			exp: false,
		},
		{
			a:   []int{1, 2},
			b:   []int{},
			exp: false,
		},
		{
			a:   []int{1},
			b:   []int{1, 2},
			exp: false,
		},
		{
			a:   []int{1, 2},
			b:   []int{1},
			exp: false,
		},
		{
			a:   []int{},
			b:   []int{1, 2, 3},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{},
			exp: false,
		},
		{
			a:   []int{1},
			b:   []int{1, 2, 3},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1},
			exp: false,
		},
		{
			a:   []int{1, 2},
			b:   []int{1, 2, 3},
			exp: false,
		},
		{
			a:   []int{1, 2, 3},
			b:   []int{1, 2},
			exp: false,
		},
	}

	for _, test := range tests {
		if rec := Equal(test.a, test.b); test.exp != rec {
			t.Fatalf("\n   given: (%d, %d)\nexpected: %t\nreceived: %t\n", test.a, test.b, test.exp, rec)
		}
	}
}

func TestFreq(t *testing.T) {
	// TODO
}

func TestJoin(t *testing.T) {
	tests := []struct {
		a   [][]int
		exp []int
	}{
		// Empty slices
		{
			a:   [][]int{},
			exp: []int{},
		},
		{
			a: [][]int{
				[]int{},
			},
			exp: []int{},
		},
		{
			a: [][]int{
				[]int{},
				[]int{},
			},
			exp: []int{},
		},
		{
			a: [][]int{
				[]int{},
				[]int{},
				[]int{},
			},
			exp: []int{},
		},

		// Single valued slices
		{
			a: [][]int{
				[]int{1},
			},
			exp: []int{1},
		},
		{
			a: [][]int{
				[]int{1},
				[]int{2},
			},
			exp: []int{1, 2},
		},
		{
			a: [][]int{
				[]int{1},
				[]int{2},
				[]int{3},
			},
			exp: []int{1, 2, 3},
		},

		// Various valued slices
		{
			a: [][]int{
				[]int{},
				[]int{},
				[]int{1, 2, 3},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{},
				[]int{1, 2, 3},
				[]int{},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{1, 2, 3},
				[]int{},
				[]int{},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{},
				[]int{1},
				[]int{2, 3},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{1},
				[]int{},
				[]int{2, 3},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{1},
				[]int{2, 3},
				[]int{},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{},
				[]int{1, 2},
				[]int{3},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{1, 2},
				[]int{},
				[]int{3},
			},
			exp: []int{1, 2, 3},
		},
		{
			a: [][]int{
				[]int{1, 2},
				[]int{3},
				[]int{},
			},
			exp: []int{1, 2, 3},
		},
	}

	for _, test := range tests {
		if rec := Join(test.a...); !Equal(test.exp, rec) {
			t.Fatalf("\n   given: %d\nexpected: %d\nreceived: %d\n", test.a, test.exp, rec)
		}

		if rec := Join2(test.a...); !Equal(test.exp, rec) {
			t.Fatalf("\n   given: %d\nexpected: %d\nreceived: %d\n", test.a, test.exp, rec)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		a, rem, exp []int
	}{
		{
			a:   []int{},
			rem: []int{},
			exp: []int{},
		},
		{
			a:   []int{1},
			rem: []int{},
			exp: []int{1},
		},
		{
			a:   []int{1},
			rem: []int{1},
			exp: []int{},
		},
		{
			a:   []int{1, 1},
			rem: []int{1},
			exp: []int{},
		},
		{
			a:   []int{1, 1},
			rem: []int{1, 1},
			exp: []int{},
		},
		{
			a:   []int{1, 2},
			rem: []int{},
			exp: []int{1, 2},
		},
		{
			a:   []int{1, 2},
			rem: []int{2},
			exp: []int{1},
		},
		{
			a:   []int{1, 2},
			rem: []int{1, 2},
			exp: []int{},
		},
		{
			a:   []int{1, 2, 3},
			rem: []int{},
			exp: []int{1, 2, 3},
		},
		{
			a:   []int{1, 2, 3},
			rem: []int{3},
			exp: []int{1, 2},
		},
		{
			a:   []int{1, 2, 3},
			rem: []int{2, 3},
			exp: []int{1},
		},
		{
			a:   []int{1, 2, 3},
			rem: []int{1, 2, 3},
			exp: []int{},
		},
		{
			a:   []int{1, 2, 3, 3},
			rem: []int{},
			exp: []int{1, 2, 3, 3},
		},
		{
			a:   []int{1, 2, 3, 3},
			rem: []int{3},
			exp: []int{1, 2},
		},
		{
			a:   []int{1, 3, 2, 3},
			rem: []int{3},
			exp: []int{1, 2},
		},
		{
			a:   []int{3, 1, 2, 3},
			rem: []int{3},
			exp: []int{1, 2},
		},
		{
			a:   []int{1, 3, 3, 2},
			rem: []int{3},
			exp: []int{1, 2},
		},
		{
			a:   []int{3, 1, 3, 2},
			rem: []int{3},
			exp: []int{1, 2},
		},
		{
			a:   []int{3, 3, 1, 2},
			rem: []int{3},
			exp: []int{1, 2},
		},
	}

	for _, test := range tests {
		if rec := Remove(test.a, test.rem...); !Equal(test.exp, rec) {
			t.Fatalf("\n   given: %d\nexpected: %d\nreceived: %d\n", test.a, test.exp, rec)
		}
	}
}

func TestResize(t *testing.T) {
	// TODO
}

func TestSearch(t *testing.T) {
	// TODO
}

func TestSwap(t *testing.T) {
	tests := []struct {
		a, cpy, exp []int
		i, j        int
		shouldPanic bool
	}{
		{
			a:           []int{},
			exp:         []int{},
			shouldPanic: true,
		},
		{
			a:           []int{1},
			exp:         []int{1},
			shouldPanic: false,
		},
		{
			a:           []int{1, 2},
			exp:         []int{1, 2},
			shouldPanic: false,
		},
		{
			a:           []int{1, 2},
			exp:         []int{1, 2},
			i:           1,
			j:           1,
			shouldPanic: false,
		},
		{
			a:           []int{1, 2},
			exp:         []int{2, 1},
			i:           0,
			j:           1,
			shouldPanic: false,
		},
		{
			a:           []int{1, 2},
			exp:         []int{2, 1},
			i:           1,
			j:           0,
			shouldPanic: false,
		},
		{
			a:           []int{1, 2, 3},
			exp:         []int{2, 1, 3},
			i:           0,
			j:           1,
			shouldPanic: false,
		},
		{
			a:           []int{1, 2, 3},
			exp:         []int{3, 2, 1},
			i:           0,
			j:           2,
			shouldPanic: false,
		},
		{
			a:           []int{1, 2, 3},
			exp:         []int{1, 3, 2},
			i:           1,
			j:           2,
			shouldPanic: false,
		},
	}

	for _, test := range tests {
		var panicked bool
		defer func() { panicked = recover() == nil }()

		test.cpy = Copy(test.a)
		Swap(test.a, test.i, test.j)
		if test.shouldPanic != panicked {
			t.Fatalf("\n   given: (%d, %d, %d)\nexpected: panic = %t\nreceived: panic = %t\n", test.cpy, test.i, test.j, test.shouldPanic, panicked)
		}

		if !test.shouldPanic && !Equal(test.a, test.exp) {
			t.Fatalf("\n   given: (%d, %d, %d)\nexpected: %d\nreceived: %d\n", test.cpy, test.i, test.j, test.exp, test.a)
		}
	}
}

func TestUnique(t *testing.T) {
	tests := []struct {
		a, exp []int
	}{
		{
			a:   []int{},
			exp: []int{},
		},
		{
			a:   []int{1},
			exp: []int{1},
		},
		{
			a:   []int{1, 1},
			exp: []int{1},
		},
		{
			a:   []int{1, 1, 1},
			exp: []int{1},
		},
		{
			a:   []int{1, 2},
			exp: []int{1, 2},
		},
		{
			a:   []int{2, 1},
			exp: []int{2, 1},
		},
		{
			a:   []int{1, 1, 2},
			exp: []int{1, 2},
		},
		{
			a:   []int{1, 2, 1},
			exp: []int{1, 2},
		},
		{
			a:   []int{2, 1, 1},
			exp: []int{2, 1},
		},
		{
			a:   []int{1, 2, 3, 3},
			exp: []int{1, 2, 3},
		},
		{
			a:   []int{1, 3, 2, 3},
			exp: []int{1, 3, 2},
		},
		{
			a:   []int{3, 1, 2, 3},
			exp: []int{3, 1, 2},
		},
		{
			a:   []int{1, 3, 3, 2},
			exp: []int{1, 3, 2},
		},
		{
			a:   []int{3, 1, 3, 2},
			exp: []int{3, 1, 2},
		},
		{
			a:   []int{3, 3, 1, 2},
			exp: []int{3, 1, 2},
		},
	}

	for _, test := range tests {
		if rec := Unique(test.a); !Equal(test.exp, rec) {
			t.Fatalf("\n   given: %d\nexpected: %d\nreceived: %d\n", test.a, test.exp, rec)
		}
	}
}

// TODO: Benchmark Join to determine which is best
