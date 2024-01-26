package main

import (
	"euler/internal/idioms"
	"fmt"
	"slices"
	"strconv"
)

const N = 1000

var ones = map[int]string{
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var tens = map[int]string{
	1: "ten",
	2: "twenty",
	3: "thirty",
	4: "forty",
	5: "fifty",
	6: "sixty",
	7: "seventy",
	8: "eighty",
	9: "ninety",
}

var teens = map[int]string{
	1: "eleven",
	2: "twelve",
	3: "thirteen",
	4: "fourteen",
	5: "fifteen",
	6: "sixteen",
	7: "seventeen",
	8: "eighteen",
	9: "nineteen",
}

var places = []string{"hundred", "thousand"}

func toDigits(n int) []int {
	var res []int
	for _, r := range strconv.Itoa(n) {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		res = append(res, i)
	}
	return res
}

func toWords(n int) []string {
	digits := toDigits(n)
	// pass 1: handle digits
	var res []string
	for i, d := range idioms.Reverse(digits) {
		switch i % 3 {
		case 0: // ones
			res = append(res, ones[d])
		case 1: // tens
			res = append(res, tens[d])
		case 2: // hundreds
			res = append(res, ones[d])
		}
	}
	// pass 2: handle teens
	res, tmp := nil, idioms.Reverse(res)
	for i, d := range tmp {
		if d == "ten" && digits[i+1] != 0 {
			tmp[i] = teens[digits[i+1]]
			tmp[i+1] = ""
		}
		res = append(res, tmp[i])
	}
	// pass 3: handle places
	res, tmp = nil, idioms.Reverse(res)
	for i, d := range tmp {
		if d == "" {
			continue
		}
		if i/3 > 0 && !slices.Contains(res, places[i/3]) {
			res = append(res, places[i/3])
		} else if i%3 == 2 {
			res = append(res, places[0])
		}
		res = append(res, d)
	}
	// pass 4: handle ands
	res, tmp = nil, idioms.Reverse(res)
	and := false
	for _, d := range tmp {
		if slices.Contains(places, d) {
			and = true
		} else if and {
			res, and = append(res, "and"), false
		}
		res = append(res, d)
	}
	return res
}

func main() {
	sum := 0
	for i := 1; i <= N; i++ {
		words, tmp := toWords(i), sum
		for _, w := range words {
			sum += len([]rune(w))
		}
		fmt.Println(i, words, sum-tmp)
	}
	fmt.Println(sum)
}
