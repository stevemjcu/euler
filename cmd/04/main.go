package main

import (
	"euler/internal/idioms"
	"fmt"
	"math"
	"strconv"
)

// Find largest palindrome made from the product of two N-digit numbers

const N = 3

func palindromes(size int) []int {
	var l []int
	lo := int(math.Pow(10, float64(size/2+size%2-1)))
	hi := lo * 10
	for i := lo; i < hi; i++ {
		a := strconv.Itoa(i)
		b := a[:len(a)-size%2]
		var s string
		for _, r := range b {
			s = string(r) + s
		}
		n, err := strconv.Atoi(a + s)
		if err != nil {
			panic(err)
		}
		l = append(l, n)
	}
	return l
}

func main() {
	for _, p := range idioms.Reverse(palindromes(N * 2)) {
		for _, f := range idioms.Factors(p) {
			if len(strconv.Itoa(f)) == N && len(strconv.Itoa(p/f)) == N {
				fmt.Println(p)
				return
			}
		}
	}
}
