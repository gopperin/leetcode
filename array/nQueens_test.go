package array

import (
	"fmt"
	"testing"
)

func ExampleSolveNQueens() {

	ret := SolveNQueens(4)

	fmt.Println(ret)

	// Output:
	// [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
}

func BenchmarkSolveNQueens(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SolveNQueens(4)
	}

	// 473656	      2529 ns/op	     432 B/op	      28 allocs/op
}

func ExampleSolveNQueens2() {

	ret := SolveNQueens2(4)

	fmt.Println(ret)

	// Output:
	// [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
}

func BenchmarkSolveNQueens2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		SolveNQueens2(4)
	}

	// 804691	      1759 ns/op	     328 B/op	      14 allocs/op
}
