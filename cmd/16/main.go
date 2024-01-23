package main

import (
	"euler/internal/idioms"
	"fmt"
)

const N = 1000

func double(num []int) []int {
	var res []int
	carry := 0
	for _, n := range idioms.Reverse(num) {
		// expressions evaluated first
		n, carry = (n*2)%10+carry, (n*2)/10
		res = append(res, n)
	}
	if carry == 1 {
		res = append(res, carry)
	}
	return idioms.Reverse(res)
}

func main() {
	num := []int{1} // 2^0
	for i := 0; i < N; i++ {
		num = double(num)
	}
	sum := 0
	for _, n := range num {
		sum += n
	}
	fmt.Println(sum)
}
