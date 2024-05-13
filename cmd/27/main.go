package main

import (
	"fmt"
	"math"
)

const N = 1000

// isPrime returns true if x is a prime number
func isPrime(x int) bool {
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return x > 1
}

// getConsecPrimes returns the number of consecutive primes for the quadratic n^2 + an + b from n=0
func getConsecPrimes(a, b int) int {
	for n := 0; ; n++ {
		if !isPrime(n*n + a*n + b) {
			return n
		}
	}
}

func main() {
	//fmt.Println(getConsecPrimes(1, 41))
	//fmt.Println(getConsecPrimes(-79, 1601))
	top := struct {
		a int
		b int
		p int
	}{}
	for a := -N + 1; a < N; a++ {
		for b := -N; b <= N; b++ {
			p := getConsecPrimes(a, b)
			if p > top.p {
				top.a, top.b, top.p = a, b, p
			}
		}
	}
	fmt.Println(top.a, top.b, top.p)
	fmt.Println(top.a * top.b)
}
