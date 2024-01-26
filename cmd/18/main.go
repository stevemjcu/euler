package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input_large.txt
var input string

func parseGraph(input string) [][]int {
	var res [][]int
	for i, r := range strings.Split(input, "\r\n") {
		res = append(res, []int{})
		for _, c := range strings.Split(r, " ") {
			n, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			res[i] = append(res[i], n)
		}
	}
	return res
}

func findMaxPathSum(i, j int, graph, cache [][]int) int {
	if i >= len(graph) {
		return 0
	}
	res := cache[i][j]
	if res != 0 {
		return res
	}
	a := findMaxPathSum(i+1, j, graph, cache)
	b := findMaxPathSum(i+1, j+1, graph, cache)
	res = graph[i][j] + max(a, b)
	cache[i][j] = res
	return res
}

func main() {
	graph := parseGraph(input)
	var cache [][]int
	for _, r := range graph {
		cache = append(cache, make([]int, len(r)))
	}
	fmt.Println(findMaxPathSum(0, 0, graph, cache))
}
