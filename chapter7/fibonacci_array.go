//  Practice 7.3
package main

import "fmt"

func main() {
	f := [50]int{}
	length := len(f)
	f[0], f[1] = 1, 1
	for i := 2; i < length; i++ {
		f[i] = f[i - 1] + f[i - 2]
	}
	//  Print one by one
	//for i := 0; i < length; i++ {
	//	fmt.Println(f[i])
	//}
	//  Print All
	 fmt.Println(f)
}
