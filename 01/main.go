package main

// Find sum of multiples of 3 or 5 below 1000

func main() {
	sum := 0
	for n := 1; n < 1000; n++ {
		if n%3 == 0 || n%5 == 0 {
			sum += n
		}
	}
	println(sum)
}
