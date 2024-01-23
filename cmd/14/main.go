package main

import "fmt"

// Find the starting number less than N which has the longest Collatz sequence

const N = 1e6

func getCollatzSequenceLength(n int) int {
	// optimization: cache in->out
	if n == 1 {
		return 1 // base case
	}
	if n%2 == 0 {
		n /= 2
	} else {
		n = 3*n + 1
	}
	return getCollatzSequenceLength(n) + 1
}

func main() {
	top, topLength := 0, 0
	for i := 1; i < N; i++ {
		length := getCollatzSequenceLength(i)
		if length > topLength {
			top = i
			topLength = length
		}
	}
	fmt.Println(top)
}
