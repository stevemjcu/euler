package main

import (
	"strconv"
	"testing"
)

func TestPrimes(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{6, 13},
		{10001, 104743},
	}
	for _, c := range cases {
		t.Run(strconv.Itoa(c.n), func(t *testing.T) {
			got := primes(c.n)[c.n-1]
			if got != c.want {
				t.Error(got, c.want)
			}
		})
	}
}
