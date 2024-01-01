package main

import (
	"euler/internal/idioms"
	"fmt"
)

// Find the difference between the sum of squares and the square of the sum of
// the first N natural numbers

const N = 100

func solve(n int) int {
	var a, b int
	for _, x := range idioms.Range(n) {
		a += x
		b += x * x
	}
	return a*a - b
}

func main() {
	fmt.Println(solve(N))
}
