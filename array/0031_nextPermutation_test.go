package array

import (
	"fmt"
	"testing"
)

func ExampleNextPermutation() {

	input := [...]int{1, 2, 3}

	NextPermutation(input[:])

	fmt.Println(input)

	// Output:
	// [1 3 2]
}

func BenchmarkNextPermutation(b *testing.B) {

	input := [...]int{1, 2, 3}

	for i := 0; i < b.N; i++ {
		NextPermutation(input[:])
	}

	// 100000000	        11.50 ns/op	       0 B/op	       0 allocs/op
}

func ExampleNextPermutation1() {

	input := [...]int{1, 2, 3}

	NextPermutation1(input[:])

	fmt.Println(input)

	// Output:
	// [1 3 2]
}

func BenchmarkNextPermutation1(b *testing.B) {

	input := [...]int{1, 2, 3}

	for i := 0; i < b.N; i++ {
		NextPermutation(input[:])
	}

	// 100000000	        10.80 ns/op	       0 B/op	       0 allocs/op
}
