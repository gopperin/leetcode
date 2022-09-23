package array

import (
	"fmt"
)

func ExamplePermute() {

	input := [...]int{1, 2, 3}

	ret := Permute(input[:])

	fmt.Println(ret)

	// Output:
	// [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
