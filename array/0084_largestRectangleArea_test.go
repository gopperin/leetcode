package array

import "fmt"

func ExampleLargestRectangleArea() {

	input := [...]int{2, 1, 5, 6, 2, 3}

	ret := LargestRectangleArea(input[:])

	fmt.Println(ret)

	// Output:
	// 10
}
