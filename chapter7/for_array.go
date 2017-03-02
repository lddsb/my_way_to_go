//  Practice 7.2
package main

import "fmt"

func main() {
	arr := [15]string{}
	a := 65
	for i := 0; i < 15; i++ {
		arr[i] = string(a + i)
	}
	fmt.Println(arr)
}
