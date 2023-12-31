package main

import "fmt"

// Find sum of even Fibonacci numbers below N

const N = 4e6

func fibonacci(n int) []int {
	var l []int
	for a, b := 1, 2; b < n; b, a = a+b, b {
		l = append(l, b)
	}
	return l
}

func solve(n int) int {
	sum := 0
	for _, f := range fibonacci(n) {
		if f%2 == 0 {
			sum += f
		}
	}
	return sum
}

func main() {
	fmt.Println(solve(N))
}
