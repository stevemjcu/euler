package main

import (
	"fmt"
	"math"
)

// Find the Nth prime number

const N = 10001

// Copied from ../03/main.go
func isPrime(n int) bool {
	for f := 2; f <= int(math.Sqrt(float64(n))); f++ {
		if n%f == 0 {
			return false
		}
	}
	return true
}

func primes(n int) []int {
	l := []int{2, 3}
	for i := 1; len(l) < n; i++ {
		a, b := 6*i-1, 6*i+1
		for _, x := range [2]int{a, b} {
			if isPrime(x) {
				l = append(l, x)
			}
		}
	}
	return l[:n]
}

func main() {
	fmt.Println(primes(N)[N-1])
}
