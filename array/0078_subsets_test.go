package array

import (
	"fmt"
)

func ExampleSubsets() {

	input := [...]int{1, 2, 3}

	ret := Subsets(input[:])

	fmt.Println(ret)

	// Output:
	// [[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]]
}

func ExampleSubsets4() {

	input := [...]int{1, 2, 3}

	ret := Subsets4(input[:])

	fmt.Println(ret)

	// Output:
	// [[] [1] [1 2] [1 2 3] [1 3] [2] [2 3] [3]]
}
