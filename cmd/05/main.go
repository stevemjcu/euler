package main

import (
	"fmt"
	"math"
	"slices"
)

// Find smallest positive number whose factors include 1..N

const N = 20

// factors returns every factor of x in ascending order
func factors(x int) []int {
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

// unroll returns a slice representing [1..x]
func unroll(x int) []int {
	var l []int
	for i := 1; i <= x; i++ {
		l = append(l, i)
	}
	return l
}

func main() {
	res := 1
	for i := 1; i <= N; i++ {
		for j := res; ; j += res {
			want, got := unroll(i), factors(j)
			if i <= len(got) && slices.Compare(want, got[:i]) == 0 {
				res = j
				break
			}
		}
	}
	fmt.Println(res)
}
