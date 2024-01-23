package main

import "fmt"

// Find the number of possible right/down routes through an NxN grid

const N = 20

func getPaths(x, y, n int) int {
	var a, b int
	if x < n {
		a = getPaths(x+1, y, n)
	}
	if y < n {
		b = getPaths(x, y+1, n)
	}
	sum := a + b
	if sum == 0 {
		return 1 // base case
	}
	return sum
}

func main() {
	fmt.Println(getPaths(0, 0, N))
}
