package main

import "fmt"

// Find sum of even Fibonacci numbers below N

const N = 4e6

func main() {
	sum := 0
	for a, b := 1, 2; b < N; b, a = a+b, b {
		if b%2 == 0 {
			sum += b
		}
	}
	fmt.Println(sum)
}
