package main

import (
	"euler/internal/idioms"
	"fmt"
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
	4: "fourty",
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
	var res []string
	for i, d := range idioms.Reverse(toDigits(n)) {
		if i%3 == 0 && i/3 > 0 {
			res = append(res, places[i/3])
		}
		switch i % 3 {
		case 0: // ones
			res = append(res, ones[d])
		case 1: // tens
			res = append(res, tens[d])
		case 2: // hundreds
			res = append(res, places[0], ones[d])
		}
	}
	return idioms.Reverse(res)
	//res, tmp := nil, idioms.Reverse(res)
	//and := false
	//for _, s := range tmp[:len(tmp)-1] {
	//	// skip empty elements
	//	if s == "" {
	//		continue
	//	}
	//	// carry the and
	//	if slices.Contains(places, s) {
	//		and = true
	//	} else if and {
	//		res = append(res, "and")
	//		and = false
	//	}
	//	res = append(res, s)
	//}
	//return res
}

func main() {
	//for i := 1; i <= N; i++ {
	//	fmt.Println(toWords(i))
	//}
	fmt.Println(toWords(101))
}
