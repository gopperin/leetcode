package array

import "fmt"

func ExampleMergeSortArray() {

	input1 := [...]int{1, 2, 3, 0, 0, 0}
	input2 := [...]int{2, 5, 6}

	MergeSortArray(input1[:], 3, input2[:], 3)

	fmt.Println(input1)

	// Output:
	// [1 2 2 3 5 6]
}
