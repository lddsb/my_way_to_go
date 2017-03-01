//  Practice 6.5
package main

import "fmt"

func main() {
	resPrint(1)
}

func resPrint(n int) {
	if n < 10 {
		resPrint(n + 1)
	}
	fmt.Println(n)
}
