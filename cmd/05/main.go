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
			want, got := idioms.Range(i), idioms.Factors(j)
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
