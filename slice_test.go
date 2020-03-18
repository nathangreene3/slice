package slice

import "testing"

// Factor a positive integer.
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
	t.Fatal(Primes(40))
}

func TestRemoveAll(t *testing.T) {
	t.Fatal(RemoveAll([]int{1, 2, 2, 2, 3, 2, 2, 4, 5, 2}, 2))
}
