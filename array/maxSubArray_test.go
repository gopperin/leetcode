package array

import (
	"fmt"
)

func ExampleMaxSubArray1() {

	input := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	ret := MaxSubArray1(input[:])

	fmt.Println(ret)

	// Output:
	// 6
}
