package array

import (
	"fmt"

	"github.com/gopperin/leetcode/structures"
)

func ExampleSortedArrayToBST() {

	values := []int{-10, -3, 0, 5, 9}

	ret := SortedArrayToBST(values[:])

	fmt.Println(structures.Tree2ints(ret))

	// Output:
	// [0 -3 9 -10 -9223372036854775808 5]
}
