package array

import (
	"fmt"
)

func ExampleCanJump() {

	input := [...]int{3, 2, 1, 0, 4}

	ret := CanJump(input[:])

	fmt.Println(ret)

	// Output:
	// false
}
