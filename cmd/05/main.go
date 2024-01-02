package main

import (
	"euler/internal/idioms"
	"fmt"
	"slices"
)

// Find smallest positive number whose factors include 1..N

const N = 20

func solve(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		for j := res; ; j += res {
			got := idioms.Factors(j)
			var want []int
			for k := 1; k <= i; k++ {
				want = append(want, k)
			}
			if i <= len(got) && slices.Compare(want, got[:i]) == 0 {
				res = j
				break
			}
		}
	}
	return res
}

func main() {
	fmt.Println(solve(N))
}
