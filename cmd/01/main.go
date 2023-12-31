package main

import "fmt"

// Find sum of multiples of 3 or 5 below N

const N = 1000

func solve(n int) int {
	sum := 0
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(solve(N))
}
