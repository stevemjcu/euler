package main

import (
	"fmt"
	"math"
)

// Copied from ../03/main.go
func isPrime(n int) bool {
	for f := 2; f <= int(math.Sqrt(float64(n))); f++ {
		if n%f == 0 {
			return false
		}
	}
	return true
}

// Copied and modified from ../07/main.go
func primes(n int) []int {
	l := []int{2, 3}
	for i := 1; ; i++ {
		a, b := 6*i-1, 6*i+1
		for _, x := range [2]int{a, b} {
			if x > n {
				return l
			} else if isPrime(x) {
				l = append(l, x)
			}
		}
	}
}

func solve(n int) int {
	sum := 0
	for _, p := range primes(n) {
		sum += p
	}
	return sum
}

func main() {
	fmt.Println(solve(2e6))
}
