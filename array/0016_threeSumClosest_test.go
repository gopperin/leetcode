package array

import (
	"fmt"
	"testing"
)

func ExampleThreeSumClosest() {

	input := [...]int{-1, 2, 1, -4}

	target := 1

	ret := ThreeSumClosest(input[:], target)

	fmt.Println(ret)

	// Output:
	// 2
}

func BenchmarkThreeSumClosest(b *testing.B) {

	input := [...]int{-1, 2, 1, -4}

	target := 1

	for i := 0; i < b.N; i++ {
		ThreeSumClosest(input[:], target)
	}

	// 8285090	       162.6 ns/op	      24 B/op	       1 allocs/op
}

func ExampleThreeSumClosest1() {

	input := [...]int{-1, 2, 1, -4}

	target := 1

	ret := ThreeSumClosest1(input[:], target)

	fmt.Println(ret)

	// Output:
	// 2
}

func BenchmarkThreeSumClosest1(b *testing.B) {

	input := [...]int{-1, 2, 1, -4}

	target := 1

	for i := 0; i < b.N; i++ {
		ThreeSumClosest1(input[:], target)
	}

	// 25321764	        48.70 ns/op	       0 B/op	       0 allocs/op
}
