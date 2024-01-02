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
		{13195, 29},
		{600851475143, 6857},
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
