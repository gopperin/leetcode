package array

import (
	"fmt"
)

func ExampleFirstMissingPositive() {

	input := [...]int{3, 4, -1, 1}

	ret := FirstMissingPositive(input[:])

	fmt.Println(ret)

	// Output:
	// 2
}
