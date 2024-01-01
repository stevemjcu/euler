package main

import (
	"euler/internal/idioms"
	"fmt"
	"math"
	"strconv"
)

// Find largest palindrome made from the product of two N-digit numbers

const N = 3

func palindromes(n int) []int {
	var l []int
	if n == 1 {
		l = append(l, 0)
	}
	lo := int(math.Pow(10, float64(n/2+n%2-1)))
	hi := lo * 10
	for i := lo; i < hi; i++ {
		a := strconv.Itoa(i)
		b := a[:len(a)-n%2]
		r := string(idioms.Reverse([]rune(b)))
		p, err := strconv.Atoi(a + r)
		if err != nil {
			panic(err)
		}
		l = append(l, p)
	}
	return l
}

func solve(n int) int {
	for _, p := range idioms.Reverse(palindromes(n * 2)) {
		for _, f := range idioms.Factors(p) {
			if len(strconv.Itoa(f)) == n && len(strconv.Itoa(p/f)) == n {
				return p
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(solve(N))
}
