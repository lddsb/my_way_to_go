//  Practice 7.1
package main

import "fmt"

func main() {
	arr1 := [10]int{}
	for i := 0; i < 10; i++ {
		arr1[i] = i * i;
	}
	arr2 := arr1
	arr2[1] = 10
	fmt.Println(arr1, arr2)
}
