//  Practice 6.1
package main

import (
	"fmt"
)

func main() {
	fmt.Print("This is no named return variables:")
	fmt.Println(multVal1(1,3))
	fmt.Print("This is named return variables:")
	fmt.Println(multVal2(3,4))
}

func multVal1(a, b int) (int, int, int) {
	sum := a + b
	product := a * b
	differ := a - b
	return sum, product, differ
}

func multVal2(a, b int) (sum int, product int, differ int) {
	sum = a + b
	product = a * b
	differ = a - b
	return
}