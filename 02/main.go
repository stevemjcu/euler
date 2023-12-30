package main

// Find sum of even Fibonacci numbers below four million

func main() {
	sum := 0
	a, b := 1, 2
	for {
		if b >= 4e6 {
			break
		}
		if b%2 == 0 {
			sum += b
		}
		b, a = a+b, b
	}
	println(sum)
}
