package main

import (
	"fmt"
)

// Find the difference between the sum of squares and the square of the sum of
// the first N natural numbers

const N = 100

func solve(n int) int {
	var a, b int
	for i := 1; i <= n; i++ {
		a += i
		b += i * i
	}
	return a*a - b
}

func main() {
	fmt.Println(solve(N))
}
