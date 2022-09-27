package array

import "fmt"

func ExampleSubsetsWithDup() {

	input := [...]int{1, 2, 2}

	ret := SubsetsWithDup(input[:])

	fmt.Println(ret)

	// Output:
	// [[] [1] [2] [1 2] [2 2] [1 2 2]]
}
