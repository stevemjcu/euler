package main

import (
	"euler/internal/idioms"
	"fmt"
)

// Find the first triangle number which has over N factors

const N = 500

func solve(n int) int {
	for i, t := 1, 0; ; i++ {
		t += i
		if len(idioms.Factors(t)) > n {
			return t
		}
	}
}

func main() {
	fmt.Println(solve(N))
}
