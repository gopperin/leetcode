package array

import (
	"fmt"
)

func ExampleGeneratePascalTriangle() {

	input := 5

	ret := GeneratePascalTriangle(input)

	fmt.Println(ret)

	// Output:
	// [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
