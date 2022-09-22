package array

import "fmt"

func ExampleFindMedianSortedArrays() {

	nums1 := [...]int{1, 2}

	nums2 := [...]int{3, 4}

	ret := FindMedianSortedArrays(nums1[:], nums2[:])

	fmt.Println(ret)

	// Output:
	// 2.5
}
