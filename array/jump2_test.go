package array

import (
	"fmt"
)

func ExampleJump2() {

	input := [...]int{2, 3, 1, 1, 4}

	ret := Jump2(input[:])

	fmt.Println(ret)

	// Output:
	// 2
}
