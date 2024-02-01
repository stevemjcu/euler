package main

import (
	"fmt"
	"math/big"
	"slices"
)

const N = 1000

func longAdd(a, b []int) []int {
	c := make([]int, 0, len(a)+1)
	if len(a) < len(b) {
		a, b = b, a
	}
	carry := 0
	for i := range a {
		x, y := a[len(a)-1-i], 0
		if i < len(b) {
			y = b[len(b)-1-i]
		}
		value := x + y + carry
		carry = value / 10
		c = append(c, value%10)
	}
	if carry != 0 {
		c = append(c, carry)
	}
	slices.Reverse(c)
	return c
}

func solveCustom() {
	a, b, i := []int{1}, []int{1}, 2
	for ; len(b) < N; i++ {
		a, b = b, longAdd(a, b)
	}
	fmt.Println(i)
}

func solveBuiltIn() {
	a, b, c := big.NewInt(1), big.NewInt(1), &big.Int{}
	limit := &big.Int{}
	limit.Exp(big.NewInt(10), big.NewInt(N-1), nil)
	var i int
	for i = 2; b.Cmp(limit) < 0; i++ {
		a, b, c = b, c.Add(a, b), a
	}
	fmt.Println(i)
}

func main() {
	solveCustom()
	solveBuiltIn()
}
