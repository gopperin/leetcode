package array

import (
	"fmt"
)

func ExampleCombinationSum() {

	input := [...]int{2, 3, 5}

	target := 8

	ret := CombinationSum(input[:], target)

	fmt.Println(ret)

	// Output:
	// [[2 2 2 2] [2 3 3] [3 5]]
}
