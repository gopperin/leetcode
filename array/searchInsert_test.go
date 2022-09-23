package array

import (
	"fmt"
)

func ExampleSearchInsert() {

	input := [...]int{1, 3, 5, 6}

	target := 5

	ret := SearchInsert(input[:], target)

	fmt.Println(ret)

	// Output:
	// 2
}
