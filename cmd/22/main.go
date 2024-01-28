package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func mergeSort[T any](s []T, less func(T, T) bool) []T {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	var a, b []T
	x := s[0]
	// sort into buckets
	for _, e := range s[1:] {
		if less(e, x) {
			a = append(a, e)
		} else {
			b = append(b, e)
		}
	}
	// sort each bucket
	a = mergeSort(a, less)
	b = mergeSort(b, less)
	return append(append(a, x), b...)
}

func lessInt(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func lessString(a, b string) bool {
	for i := 0; i < min(len(a), len(b)); i++ {
		if !(a[i] == b[i]) {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func parseInput(s string) []string {
	return strings.Split(s[1:len(s)-1], `","`)
}

func main() {
	sum := 0
	for i, name := range mergeSort(parseInput(input), lessString) {
		score := 0
		for _, r := range name {
			score += int(r - 'A' + 1)
		}
		sum += score * (i + 1)
		//if name == "COLIN" {
		//	fmt.Println(i+1, name, score * (i + 1))
		//}
	}
	fmt.Println(sum)
}
