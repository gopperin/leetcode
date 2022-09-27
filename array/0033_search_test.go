package array

import (
	"fmt"
)

func ExampleSearch33() {

	input := [...]int{4, 5, 6, 7, 0, 1, 2}

	target := 0

	ret := Search33(input[:], target)

	fmt.Println(ret)

	// Output:
	// 4
}
