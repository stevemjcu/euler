package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestFibonacci(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{1, []int{}},
		{2, []int{1}},
		{10, []int{1, 2, 3, 5, 8}},
		{100, []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89}},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			got := fibonacci(c.n)
			if slices.Compare(got, c.want) != 0 {
				t.Error(got, c.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{100, 44},
		{4e6, 4613732},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			got := solve(c.n)
			if got != c.want {
				t.Error(got, c.want)
			}
		})
	}
}
