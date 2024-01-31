package main

import (
	"fmt"
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

func main() {
	a, b, i := []int{1}, []int{1}, 2
	for ; len(b) < N; i++ {
		a, b = b, longAdd(a, b)
	}
	fmt.Println(i)
}
