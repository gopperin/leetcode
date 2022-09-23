package array

import (
	"fmt"
)

func ExamplePermuteUnique() {

	input := [...]int{1, 1, 2}

	ret := PermuteUnique(input[:])

	fmt.Println(ret)

	// Output:
	// [[1 1 2] [1 2 1] [2 1 1]]
}
