package main

import (
	"euler/internal/idioms"
	"fmt"
)

// Find the difference between the sum of squares and the square of the sum of
// the first N natural numbers

const N = 100

func solve(n int) int {
	l := idioms.Range(n)
	a := idioms.Reduce(l, func(e int) int { return e })
	b := idioms.Reduce(l, func(e int) int { return e * e })
	return a*a - b
}

func main() {
	fmt.Println(solve(N))
}
