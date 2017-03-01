//  Practice 6.8
package main

import "fmt"

func main() {
	fv := func() {fmt.Println("Hello World!")}
	fv()
	fmt.Println(f())
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
