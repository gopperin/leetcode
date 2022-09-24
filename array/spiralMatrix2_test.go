package array

import (
	"fmt"
)

func ExampleGenerateMatrix() {

	ret := GenerateMatrix(3)

	fmt.Println(ret)

	// Output:
	// [[1 2 3] [8 9 4] [7 6 5]]
}
