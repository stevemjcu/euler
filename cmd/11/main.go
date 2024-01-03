package main

import (
	_ "embed"
	"euler/internal/idioms"
	"fmt"
	"strconv"
	"strings"
)

// Find greatest product of N adjacent numbers in the grid

const N = 4

//go:embed input.txt
var input string

type vec struct {
	x, y int
}

func parse(in string) [][]int {
	var out [][]int
	for i, row := range strings.Split(in, "\r\n") {
		out = append(out, []int{})
		for _, ele := range strings.Split(row, " ") {
			x, err := strconv.Atoi(ele)
			if err != nil {
				panic(err)
			}
			out[i] = append(out[i], x)
		}
	}
	return out
}

func reduce(n int, m [][]int, dir vec) int {
	top := 0
	for i := 0; i < len(m)-(n-1)*dir.y; i++ {
		for j := 0; j < len(m)-(n-1)*dir.x; j++ {
			acc := 1
			for k := 0; k < n; k++ {
				acc *= m[i+k*dir.y][j+k*dir.x]
			}
			top = max(top, acc)
		}
	}
	return top
}

func solve(n int, m [][]int) int {
	top := 0
	top = max(top, reduce(n, m, vec{0, 1}))
	top = max(top, reduce(n, m, vec{1, 0}))
	top = max(top, reduce(n, m, vec{1, 1}))
	top = max(top, reduce(n, idioms.Reverse(m), vec{1, 1}))
	return top
}

func main() {
	fmt.Println(solve(N, parse(input)))
}
