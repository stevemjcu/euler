package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// Find the first N digits of a large sum

const N = 10

//go:embed input.txt
var input string

func parse(in string) []float64 {
	var out []float64
	for _, r := range strings.Split(in, "\r\n") {
		f, err := strconv.ParseFloat(r, 64)
		if err != nil {
			panic(err)
		}
		out = append(out, f)
	}
	return out
}

func solve(n int, l []float64) int {
	var sum float64
	for _, f := range l {
		sum += f
	}
	s := strconv.FormatFloat(sum, 'g', -1, 64)
	s = strings.Replace(s, ".", "", 1)
	i, err := strconv.Atoi(string([]rune(s)[:n]))
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	fmt.Println(solve(N, parse(input)))
}
