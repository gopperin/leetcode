package array

import (
	"fmt"
)

func ExampleSearchRange() {

	input := [...]int{5, 7, 7, 8, 8, 10}

	target := 8

	ret := SearchRange(input[:], target)

	fmt.Println(ret)

	// Output:
	// [3 4]
}
