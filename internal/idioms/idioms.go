package idioms

import (
	"math"
	"slices"
)

type Number interface {
	int | int64 | float64
}

// Factors returns every factor of x in ascending order.
func Factors(x int) []int {
	var a, b []int
	for i := 1; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			a = append(a, i)
			if i != x/i {
				b = append(b, x/i)
			}
		}
	}
	slices.Reverse(b)
	return append(a, b...)
}

// Range returns the slice representing [1..x].
func Range(x int) []int {
	var l []int
	for i := 1; i <= x; i++ {
		l = append(l, i)
	}
	return l
}

// Reverse returns the reverse of the slice s.
func Reverse[S ~[]E, E any](s S) S {
	r := slices.Clone(s)
	slices.Reverse(r)
	return r
}

// Reduce applies an accumulative function to each element of the slice s.
func Reduce[S ~[]E, E any, R Number](s S, f func(E) R) R {
	var r R
	for _, e := range s {
		r += f(e)
	}
	return r
}

// Sum returns the sum of each element of the slice s.
func Sum[S ~[]E, E Number](s S) E {
	return Reduce(s, func(e E) E {
		return e
	})
}
