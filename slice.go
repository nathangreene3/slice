package slice

func Contains[T comparable](s []T, v T) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return true
		}
	}

	return false
}

func ContainsAny[T comparable](s []T, vs ...T) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(vs); j++ {
			if s[i] == vs[j] {
				return true
			}
		}
	}

	return false
}

func Copy[T any](s []T) []T {
	return append(make([]T, 0, len(s)), s...)
}

func Count[T comparable](s []T, v T) int {
	var n int
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			n++
		}
	}

	return n
}

func Equal[T comparable](ss ...[]T) bool {
	for i := 1; i < len(ss); i++ {
		if len(ss[0]) != len(ss[i]) {
			return false
		}

		for j := 0; j < len(ss[0]); j++ {
			if ss[0][j] != ss[i][j] {
				return false
			}
		}
	}

	return true
}

func Filter[T any](s []T, f func(T) bool) []T {
	filtered := make([]T, 0, len(s))
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			filtered = append(filtered, s[i])
		}
	}

	return filtered
}

func Fold[T0, T1 any](s []T0, seed T1, f func(T1, T0) T1) T1 {
	for i := 0; i < len(s); i++ {
		seed = f(seed, s[i])
	}

	return seed
}

func Freq[T comparable](s []T) map[T]int {
	freq := make(map[T]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	return freq
}

func Generate[T any](n int, f func(int) T) []T {
	s := make([]T, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, f(i))
	}

	return s
}

func Join[T any](ss ...[]T) []T {
	var n int
	for i := 0; i < len(ss); i++ {
		n += len(ss[i])
	}

	joined := make([]T, 0, n)
	for i := 0; i < len(ss); i++ {
		joined = append(joined, ss[i]...)
	}

	return joined
}

func Map[T0, T1 any](s []T0, f func(T0) T1) []T1 {
	t := make([]T1, 0, len(s))
	for i := 0; i < len(s); i++ {
		t = append(t, f(s[i]))
	}

	return t
}

func Reduce[T any](s []T, f func(_, _ T) T) T {
	var r T
	if len(s) != 0 {
		r = s[0]
	}

	return Fold(s, r, f)
}

func Remove[T comparable](s []T, vs ...T) []T {
	return Filter(s, func(v T) bool { return !Contains(s, v) })
}

func Resize[T any](s []T, n, c int) []T {
	if len(s) < n {
		n = len(s)
	}

	return append(make([]T, 0, c), s[:n]...)
}

func Search[T comparable](s []T, v any) int {
	var i int
	for ; i < len(s) && s[i] != v; i++ {
	}

	return i
}

func SubSlice[T any](s []T, i, j int) []T {
	return append(make([]T, 0, j-i), s[i:j]...)
}

func Swap[T any](s []T, i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Unique[T comparable](ss ...[]T) []T {
	var n int
	for i := 0; i < len(ss); i++ {
		n += len(ss[i])
	}

	var unique []T
	if n != 0 {
		observed := make(map[T]struct{})
		for i := 0; i < len(ss); i++ {
			for j := 0; j < len(ss[i]); j++ {
				if _, ok := observed[ss[i][j]]; !ok {
					unique = append(unique, ss[i][j])
					observed[ss[i][j]] = struct{}{}
				}
			}
		}
	}

	return unique
}
