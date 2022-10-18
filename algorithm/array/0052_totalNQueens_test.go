package array

import (
	"fmt"
)

type question52 struct {
	para52
	ans52
}

// para 是参数
// one 代表第一个参数
type para52 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans52 struct {
	one int
}

func ExampleTotalNQueens() {

	qs := []question52{

		{
			para52{1},
			ans52{1},
		},

		{
			para52{2},
			ans52{0},
		},
		{
			para52{3},
			ans52{0},
		},

		{
			para52{4},
			ans52{2},
		},

		{
			para52{5},
			ans52{10},
		},

		{
			para52{6},
			ans52{4},
		},

		{
			para52{7},
			ans52{40},
		},

		{
			para52{8},
			ans52{92},
		},

		{
			para52{9},
			ans52{352},
		},

		{
			para52{10},
			ans52{724},
		},
	}

	for _, q := range qs {
		_, p := q.ans52, q.para52
		fmt.Println(p, TotalNQueens(p.one))
	}
	// output:
	// {1} 1
	// {2} 0
	// {3} 0
	// {4} 2
	// {5} 10
	// {6} 4
	// {7} 40
	// {8} 92
	// {9} 352
	// {10} 724
}
