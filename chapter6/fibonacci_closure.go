// Practice 6.9
package main

import "fmt"

func main() {
	f := closureFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("f(", i, "):", f())
	}
}

func closureFibonacci() func() int {
	f0, f1, sum := -1, 1, 0
	return func() int {
		if f0 == -1 {
			f0++
			return f1
		} else {
			sum = f0 + f1
			f0 = f1
			f1 = sum
			return sum
		}
	}
}