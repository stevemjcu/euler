package main

import (
	"fmt"
	"math/big"
)

const N = 100

func main() {
	m := make(map[string]bool)
	for a := int64(2); a <= N; a++ {
		for b := int64(2); b <= N; b++ {
			c := new(big.Int)
			c.Exp(big.NewInt(a), big.NewInt(b), nil)
			m[c.String()] = true
		}
	}
	fmt.Println(len(m))
}
