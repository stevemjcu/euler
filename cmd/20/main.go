package main

import "fmt"

func toDigits(n int) []int {
	var res []int
	for ; n != 0; n /= 10 {
		res = append([]int{n % 10}, res...)
	}
	return res
}

func longAdd(a, b []int) []int {
	if len(a) < len(b) {
		a, b = b, a
	}
	z := make([]int, len(a)-len(b))
	b = append(z, b...) // pad with 0
	var c []int
	carry := 0
	for i := len(a) - 1; i >= 0; i-- {
		x := a[i] + b[i] + carry
		carry, x = x/10, x%10
		c = append([]int{x}, c...)
	}
	if carry != 0 {
		c = append([]int{carry}, c...)
	}
	return c
}

func longMultiply(a, b []int) []int {
	var d []int
	for i := len(b) - 1; i >= 0; i-- {
		c := make([]int, len(b)-1-i)
		carry := 0
		for j := len(a) - 1; j >= 0; j-- {
			x := a[j]*b[i] + carry
			carry, x = x/10, x%10
			c = append([]int{x}, c...)
		}
		if carry != 0 {
			c = append([]int{carry}, c...)
		}
		d = longAdd(c, d)
	}
	return d
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialBig(n int) []int {
	if n == 1 {
		return []int{1}
	}
	a, b := toDigits(n), factorialBig(n-1)
	return longMultiply(a, b)
}

func main() {
	sum := 0
	for _, x := range factorialBig(100) {
		sum += x
	}
	fmt.Println(sum)
}
