package main

import (
	"fmt"
	"slices"
)

const N = 1e6

var input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func permutations(s []int) [][]int {
	if len(s) == 1 {
		return [][]int{s}
	}
	var res [][]int
	for i, x := range s {
		// append can reuse the slice, so must clone
		t := slices.Delete(slices.Clone(s), i, i+1)
		for _, p := range permutations(t) {
			res = append(res, append([]int{x}, p...))
		}
	}
	return res
}

func compareLexicographic(a, b []int) int {
	for i := range a {
		if a[i] != b[i] {
			return a[i] - b[i]
		}
	}
	return len(a) - len(b)
}

func main() {
	fmt.Println(permutations(input)[N-1])
	//fmt.Println(slices.IsSortedFunc(permutations(input), compareLexicographic))
}
