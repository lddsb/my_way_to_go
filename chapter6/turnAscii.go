//  Practice 6.7
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abc我efg反阿兰ss"
	//str := "ab cdefgh"
	str2 := strings.Split(str, "")
LABEL1:
	index := strings.IndexFunc(str, IsAscii)
	//fmt.Print(index)
	//return
	if -1 != index {
		fmt.Println(index, str2[index])
		str2[index] = "?"
		str = strings.Join(str2, "")
		goto LABEL1
	}
	fmt.Println("Converted string:", str)
	fmt.Println("Converted string slice:", str2)
}

func IsAscii(c rune) bool {
	if c > 255 {
		return true
	}
	return false
}
