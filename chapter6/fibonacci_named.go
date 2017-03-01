//  Practice 6.4
package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 10; i++ {
		index, result := fibonacci(i)
		fmt.Printf("fibonacci(%d) return index: %d, res: %d", i, index, result)
		fmt.Println()
	}
	fmt.Println()
}

func fibonacci(n int) (index, res int) {
	if n <= 1 {
		index = n
		res = 1
	} else {
		_, res1 := fibonacci(n - 1)
		_, res2 := fibonacci(n - 2)
		index = n
		res = res1 + res2
	}
	return
}
