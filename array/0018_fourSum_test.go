package array

import (
	"fmt"
	"testing"
)

func ExampleFourSum() {

	input := [...]int{1, 0, -1, 0, -2, 2}

	target := 0

	ret := FourSum(input[:], target)

	fmt.Println(ret)

	// Output:
	// [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
}

func BenchmarkFourSum(b *testing.B) {

	input := [...]int{1, 0, -1, 0, -2, 2}

	target := 0

	for i := 0; i < b.N; i++ {
		FourSum(input[:], target)
	}

	// 1555382	       801.9 ns/op	     288 B/op	       7 allocs/op
}

func ExampleFourSum1() {

	input := [...]int{1, 0, -1, 0, -2, 2}

	target := 0

	ret := FourSum1(input[:], target)

	fmt.Println(ret)

	// Output:
	// [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
}

func BenchmarkFourSum1(b *testing.B) {

	input := [...]int{1, 0, -1, 0, -2, 2}

	target := 0

	for i := 0; i < b.N; i++ {
		FourSum1(input[:], target)
	}

	// 554988	      2541 ns/op	     664 B/op	      35 allocs/op
}
