package array

import (
	"fmt"
)

func ExampleSortColors() {

	input := [...]int{2, 0, 2, 1, 1, 0}

	SortColors(input[:])

	fmt.Println(input)

	// Output:
	// [0 0 1 1 2 2]
}
