package main

import "math"

// Find largest prime factor of 600851475143

func isPrime(n int) bool {
	for f := 2; f <= int(math.Sqrt(float64(n))); f++ {
		if n%f == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 600851475143
	for f := 1; f <= int(math.Sqrt(float64(n))); f++ {
		if n%f == 0 && isPrime(n/f) {
			println(n / f)
			return
		}
	}
	for f := int(math.Sqrt(float64(n))); f >= 1; f-- {
		if n%f == 0 && isPrime(f) {
			println(f)
			return
		}
	}
}
