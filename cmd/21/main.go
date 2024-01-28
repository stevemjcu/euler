package main

import (
	"fmt"
	"math"
)

const N = 10000

func properDivisors(n int) []int {
	var a, b []int
	sqrt := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if n%i == 0 {
			a = append(a, i)
			if i != 1 && i != sqrt {
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
		for _, x := range properDivisors(i) {
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
