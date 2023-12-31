package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

// Find largest palindrome made from the product of two N-digit numbers

const N = 3

// palindromes returns every palindrome with the length z
func palindromes(z int) []int {
	var l []int
	lo := int(math.Pow(10, float64(z/2+z%2-1)))
	hi := lo * 10
	for i := lo; i < hi; i++ {
		a := strconv.Itoa(i)
		b := a[:len(a)-z%2]
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
	slices.Reverse(l)
	return l
}

// factors returns every factor of x
func factors(x int) [][2]int {
	var l [][2]int
	for i := 1; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			l = append(l, [2]int{i, x / i})
		}
	}
	return l
}

func main() {
	for _, p := range palindromes(N * 2) {
		for _, f := range factors(p) {
			if len(strconv.Itoa(f[0])) == N && len(strconv.Itoa(f[1])) == N {
				fmt.Println(p)
				return
			}
		}
	}
}
