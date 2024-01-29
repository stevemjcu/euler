package main

import (
	"fmt"
	"math"
)

const N = 28123

// Copied from cmd/21/main.go
func divisors(n int) []int {
	var a, b []int
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			a = append(a, i)
			if n/i != i {
				b = append([]int{n / i}, b...)
			}
		}
	}
	return append(a, b...)
}

func isAbundant(n int) bool {
	sum := 0
	s := divisors(n)
	for _, x := range s[:len(s)-1] {
		sum += x
	}
	return sum > n
}

func main() {
	// enumerate abundant numbers
	var nums []int
	for i := 1; i <= N; i++ {
		if isAbundant(i) {
			nums = append(nums, i)
		}
	}
	// enumerate abundant sums
	sums := make(map[int]bool)
	for _, x := range nums {
		for _, y := range nums {
			sums[x+y] = true
		}
	}
	// enumerate excluded numbers
	sum := 0
	for i := 1; i <= N; i++ {
		if !sums[i] {
			sum += i
		}
	}
	fmt.Println(sum)
}
