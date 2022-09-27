package array

import (
	"fmt"
)

func ExampleTrap() {

	input := [...]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	ret := Trap(input[:])

	fmt.Println(ret)

	// Output:
	// 6
}
