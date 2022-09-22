package array

import (
	"fmt"
)

func ExampleTwoSum() {

	nums := [...]int{2, 7, 11, 15}

	target := 22

	ret := TwoSum(nums[:], target)

	fmt.Println(ret)

	// Output:
	// [1 3]
}
