package main

import (
	"slices"
	"strconv"
	"testing"
)

func TestPalindrome(t *testing.T) {
	cases := []struct {
		n    int
		want []int
	}{
		{1, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{2, []int{11, 22, 33, 44, 55, 66, 77, 88, 99}},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			got := palindromes(c.n)
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
		{2, 9009},
		{3, 906609},
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
