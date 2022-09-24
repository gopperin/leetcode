package array

import (
	"fmt"
	"testing"
)

func ExampleMaxSubArray() {

	ret := SolveNQueens(4)

	fmt.Println(ret)

	// Output:
	// [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
}

func BenchmarkMaxSubArray(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SolveNQueens(4)
	}

	// 473656	      2529 ns/op	     432 B/op	      28 allocs/op
}

func ExampleMaxSubArray1() {

	input := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	ret := MaxSubArray1(input[:])

	fmt.Println(ret)

	// Output:
	// 6
}

func BenchmarkMaxSubArray1(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SolveNQueens2(4)
	}

	// 804691	      1759 ns/op	     328 B/op	      14 allocs/op
}
