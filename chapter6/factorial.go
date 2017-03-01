//  Practice 6.6
package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(30)
	fmt.Print(factorial(n))
}

func factorial(n *big.Int) (sum *big.Int) {
	zero := big.NewInt(0)
	if n.Cmp(zero) <= 0 {
		sum = big.NewInt(1)
	} else {
		sum = new(big.Int)
		sum.Set(n)
		sum.Mul(sum, factorial(n.Sub(n, big.NewInt(1))))
	}
	return
}
