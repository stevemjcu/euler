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
		{4, 5832},
		{13, 23514624000},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			got := solve(c.n, parse(input))
			if got != c.want {
				t.Error(got, c.want)
			}
		})
	}
}
