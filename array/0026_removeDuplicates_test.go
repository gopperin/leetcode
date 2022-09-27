package array

import (
	"fmt"
	"testing"
)

func ExampleRemoveDuplicates() {

	input := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	ret := RemoveDuplicates(input[:])

	fmt.Println(ret)
	fmt.Println(input[:ret])

	// Output:
	// 5
	// [0 1 2 3 4]
}

func BenchmarkRemoveDuplicates(b *testing.B) {

	input := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	for i := 0; i < b.N; i++ {
		RemoveDuplicates(input[:])
	}

	// 34406098	        33.65 ns/op	       0 B/op	       0 allocs/op
}

func ExampleRemoveDuplicates1() {

	input := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	ret := RemoveDuplicates1(input[:])

	fmt.Println(ret)
	fmt.Println(input[:ret])

	// Output:
	// 5
	// [0 1 2 3 4]
}

func BenchmarkRemoveDuplicates1(b *testing.B) {

	input := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	for i := 0; i < b.N; i++ {
		RemoveDuplicates1(input[:])
	}

	// 117504350	         8.882 ns/op	       0 B/op	       0 allocs/op
}
