package main

import (
	"fmt"
	"math"
)

// Find the product of the Pythagorean triplet whose sum is N

const N = 1000

func solve(n int) int {
	for a := 1; a < n-1; a++ {
		for b := a + 1; b < n; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if c == float64(int(c)) && a+b+int(c) == n {
				return a * b * int(c)
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(solve(N))
}
