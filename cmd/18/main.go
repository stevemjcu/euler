package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
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

func findMaxPathSum(i, j int, graph [][]int) int {
	if i >= len(graph) {
		return 0
	}
	a := findMaxPathSum(i+1, j, graph)
	b := findMaxPathSum(i+1, j+1, graph)
	return graph[i][j] + max(a, b)
}

func main() {
	fmt.Println(findMaxPathSum(0, 0, parseGraph(input)))
}
