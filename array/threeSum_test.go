package array

import (
	"fmt"
	"testing"
)

func ExampleThreeSum() {

	input := [...]int{-1, 0, 1, 2, -1, -4}

	ret := ThreeSum(input[:])

	fmt.Println(ret)

	// Output:
	// [[-1 -1 2] [-1 0 1]]
}

func BenchmarkThreeSum(b *testing.B) {

	input := [...]int{-1, 0, 1, 2, -1, -4}

	for i := 0; i < b.N; i++ {
		ThreeSum(input[:])
	}

	// 1817342	       657.3 ns/op	     144 B/op	       5 allocs/op
}

func ExampleThreeSum1() {

	input := [...]int{-1, 0, 1, 2, -1, -4}

	ret := ThreeSum1(input[:])

	fmt.Println(ret)

	// Output:
	// [[-1 0 1] [-1 -1 2]]
}

func BenchmarkThreeSum1(b *testing.B) {

	input := [...]int{-1, 0, 1, 2, -1, -4}

	for i := 0; i < b.N; i++ {
		ThreeSum1(input[:])
	}

	// 449923	      2290 ns/op	     312 B/op	      11 allocs/op
}
