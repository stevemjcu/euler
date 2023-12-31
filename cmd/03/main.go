package main

import (
	"euler/internal/idioms"
	"fmt"
	"math"
)

// Find largest prime factor of N

const N = 600851475143

func isPrime(n int) bool {
	for f := 2; f <= int(math.Sqrt(float64(n))); f++ {
		if n%f == 0 {
			return false
		}
	}
	return true
}

func main() {
	for _, f := range idioms.Reverse(idioms.Factors(N)) {
		if isPrime(f) {
			fmt.Println(f)
			return
		}
	}
}
