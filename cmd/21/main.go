package main

import (
	"fmt"
	"math"
)

const N = 10000

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

func main() {
	sums := make(map[int]int)
	for i := 1; i < N; i++ {
		sums[i] = 0
		s := divisors(i)
		for _, x := range s[:len(s)-1] {
			sums[i] += x
		}
	}
	sum := 0
	for i := 1; i < N; i++ {
		j := sums[i]
		k, ok := sums[j]
		if ok && i == k && i != j {
			fmt.Println(i, j)
			sum += i
		}
	}
	fmt.Println(sum)
}
