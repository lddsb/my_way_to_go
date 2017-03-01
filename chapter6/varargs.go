//  Practice 6.3
package main

import (
	"fmt"
)

func main() {
	varargs(1, 2, 4, 5, 6)
}

func varargs(a ... int) {
	length := len(a)
	if length > 0 {
		for i, val := range a {
			fmt.Println("The index ", i, " value is:", val)
		}
	}
}
