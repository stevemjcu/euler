package main

import (
	"fmt"
	"slices"
)

// Find the value of d < N for which 1/d contains the largest recurring cycle.

const N = 1000

// findRecurringCycleLength returns the length of the recurring cycle for 1/x.
func findRecurringCycleLength(x int) int {
	r := 1
	var s []int
	for {
		for r < x {
			r *= 10
		}
		r %= x
		if r == 0 {
			return 0 // no cycle
		}
		if slices.Contains(s, r) {
			return findRemainderDistance(r, r, x)
		}
		s = append(s, r)
	}
}

// findRemainderDistance returns the length between remainders produced via long division.
func findRemainderDistance(a, b, x int) int {
	d := 0
	for {
		for a < x {
			a *= 10
			d++
		}
		a %= x
		if a == b {
			return d
		}
	}
}

func main() {
	t, tl := 0, 0
	for i := 2; i < N; i++ {
		if il := findRecurringCycleLength(i); il > tl {
			t, tl = i, il
		}
	}
	fmt.Println(t)
}

// Not needed but saving for later
//func lengthOfLongestSubstring(s string) int {
//	var r = []rune(s)
//	var i, j int
//	top := 0
//	set := map[rune]int{}
//	for i, j = 0, 0; j < len(r); j++ {
//		k, ok := set[r[j]]
//		if ok {
//			top = max(top, j-i)
//			i, j = k+1, k+1
//			set = map[rune]int{}
//		}
//		set[r[j]] = j
//	}
//	return max(top, j-i)
//}
