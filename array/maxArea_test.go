package array

import "fmt"

func ExampleMaxArea() {

	input := [...]int64{1, 8, 6, 2, 5, 4, 8, 3, 7}

	ret := MaxArea(input[:])

	fmt.Println(ret)

	// Output:
	// 49
}
