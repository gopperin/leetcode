package array

import (
	"fmt"
)

func ExampleRemoveDuplicates2() {

	input := [...]int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	ret := RemoveDuplicates2(input[:])

	fmt.Println(ret)
	fmt.Println(input[:ret])

	// Output:
	// 7
	// [0 0 1 1 2 3 3]
}
