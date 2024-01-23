package main

import (
	"fmt"
	"strconv"
)

// Find the number of possible right/down routes through an NxN grid

const N = 20

func toKey(x, y int) string {
	return strconv.Itoa(x) + "-" + strconv.Itoa(y)
}

func getPaths(x, y int, cache map[string]int) int {
	res, ok := cache[toKey(x, y)]
	if ok {
		return res // cache hit
	}
	var a, b int
	if x > 0 {
		a = getPaths(x-1, y, cache)
	}
	if y > 0 {
		b = getPaths(x, y-1, cache)
	}
	res = a + b
	if res == 0 {
		res = 1 // base case
	}
	cache[toKey(x, y)] = res
	return res
}

func main() {
	cache := make(map[string]int)
	for i := 0; i <= N; i++ {
		fmt.Println(i, getPaths(i, i, cache))
	}
}
