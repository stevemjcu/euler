package main

import (
	"euler/internal/idioms"
	"fmt"
	"math"
)

const N = 5
const M = 1e8 // arbitrary limit

func main() {
	sum := 0
	for i := 0; i < M; i++ {
		n := 0
		for _, d := range idioms.ToDigits(i) {
			n += int(math.Pow(float64(d), N))
		}
		if i == n {
			fmt.Println(idioms.ToDigits(i))
			sum += i
		}
	}
	fmt.Println(sum - 1) //exclude 1
}
