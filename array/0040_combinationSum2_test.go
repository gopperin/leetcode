package array

import (
	"fmt"
)

func ExampleCombinationSum2() {

	input := [...]int{2, 5, 2, 1, 2}

	target := 5

	ret := CombinationSum2(input[:], target)

	fmt.Println(ret)

	// Output:
	// [[1 2 2] [5]]
}
