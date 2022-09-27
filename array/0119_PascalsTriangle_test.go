package array

import (
	"fmt"
)

func ExampleGetPascalTriangleRow() {

	input := 1

	ret := GetPascalTriangleRow(input)

	fmt.Println(ret)

	// Output:
	// [1 1]
}
