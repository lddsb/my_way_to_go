//  Practice 6.2
package main

import (
	"fmt"
	"math"
	"errors"
)

func main() {
	fmt.Print()
	ret1, err1 := errorVal1(-1)
	if err1 != nil {
		fmt.Println("Error!Return values are:", ret1, err1)
	} else {
		fmt.Println("OK！Return values are:", ret1, err1)
	}

	if ret2, err2 := errorVal2(4); err2 != nil {
		fmt.Println("Error!Return values are:", ret2, err2)
	} else {
		fmt.Println("OK！Return values are:", ret2, err2)
	}
}

func errorVal1(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("This numbers is illegal!")
	}
	return math.Sqrt(f), nil
}

func errorVal2(f float64) (ret float64, err error) {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("This numbers is illegal!")
	} else {
		ret = math.Sqrt(f)
		//err = nil
	}
	return
}
