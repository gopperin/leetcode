package array

import (
	"fmt"
)

func ExampleRemoveElement() {

	input := [...]int{0, 1, 2, 2, 3, 0, 4, 2}

	val := 2

	ret := RemoveElement(input[:], val)

	fmt.Println(ret)
	fmt.Println(input[:ret])

	// Output:
	// 5
	// [0 1 3 0 4]
}
