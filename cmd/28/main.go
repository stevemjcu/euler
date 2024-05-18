package main

import "fmt"

const N = 1001

func main() {
	n, sum := 1, 1
	// for each layer
	for i := 1; i < (N+1)/2; i++ {
		// for each hop
		for j := 0; j < 4; j++ {
			n += 2 * i
			sum += n
		}
	}
	fmt.Println(sum)
}
