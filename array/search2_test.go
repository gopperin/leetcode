package array

import (
	"fmt"
)

func ExampleSearch2() {

	values := []int{2, 5, 6, 0, 0, 1, 2}

	target := 3

	ret := Search2(values[:], target)

	fmt.Println(ret)

	// Output:
	// false
}
