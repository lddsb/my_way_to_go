//  Practice 7.1
package main

import "fmt"

func main() {
	arr1 := [10]int{}
	for i := 0; i < 10; i++ {
		arr1[i] = i * 2;
	}
	//  print arr1
	fmt.Println("This is for:")
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	}
	//  for range
	fmt.Println("This is for range:")
	for i, val := range arr1 {
		fmt.Printf("Array index %d is %d\n", i, val)
	}
	arr2 := arr1
	arr2[1] = 10
	fmt.Println(arr1, arr2)
}
