package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// Find the N adjacent digits which have the greatest product

const N = 13

//go:embed input.txt
var input string

func parse(in string) []int {
	var out []int
	in = strings.ReplaceAll(in, "\r\n", "")
	for _, r := range in {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		out = append(out, n)
	}
	return out
}

func solve(n int, l []int) int {
	top := 0
	for i := 0; i <= len(l)-n; i++ {
		acc := 1
		for _, e := range l[i : i+n] {
			acc *= e
		}
		top = max(top, acc)
	}
	return top
}

func main() {
	fmt.Println(solve(N, parse(input)))
}
