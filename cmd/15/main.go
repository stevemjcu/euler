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

func getPaths(x, y, n int, cache map[string]int) int {
	res, ok := cache[toKey(x, y)]
	if ok {
		return res // cache hit
	}
	var a, b int
	if x < n {
		a = getPaths(x+1, y, n, cache)
	}
	if y < n {
		b = getPaths(x, y+1, n, cache)
	}
	res = a + b
	if res == 0 {
		res = 1 // base case
	}
	cache[toKey(x, y)] = res
	return res
}

func main() {
	for i := 1; i <= N; i++ {
		cache := make(map[string]int)
		fmt.Println(i, getPaths(0, 0, i, cache))
	}
}
