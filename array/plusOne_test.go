package array

import (
	"fmt"
)

func ExamplePlusOne() {

	input := [...]int{4, 3, 2, 1}

	ret := PlusOne(input[:])

	fmt.Println(ret)

	// Output:
	// [4 3 2 2]
}
