package main

import (
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{10, 17},
		{2e6, 142913828922},
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
