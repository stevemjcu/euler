package main

import (
	"fmt"
	"math"
	"slices"
)

// Find largest prime factor of 600851475143

func factors(n int) []int {
	var a, b []int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			a = append(a, i)
			if i != n/i {
				b = append(b, n/i)
			}
		}
	}
	slices.Reverse(a)
	return append(b, a...)
}

func isPrime(n int) bool {
	for f := 2; f <= int(math.Sqrt(float64(n))); f++ {
		if n%f == 0 {
			return false
		}
	}
	return true
}

func main() {
	for _, f := range factors(600851475143) {
		if isPrime(f) {
			fmt.Println(f)
			return
		}
	}
}
