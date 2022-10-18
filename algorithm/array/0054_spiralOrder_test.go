package array

import (
	"fmt"
)

type question54 struct {
	para54
	ans54
}

// para 是参数
// one 代表第一个参数
type para54 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans54 struct {
	one []int
}

func ExampleSpiralOrder() {

	qs := []question54{

		{
			para54{[][]int{{3}, {2}}},
			ans54{[]int{3, 2}},
		},

		{
			para54{[][]int{{2, 3}}},
			ans54{[]int{2, 3}},
		},

		{
			para54{[][]int{{1}}},
			ans54{[]int{1}},
		},

		{
			para54{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
			ans54{[]int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		},
		{
			para54{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}},
			ans54{[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		},
	}

	for _, q := range qs {
		_, p := q.ans54, q.para54
		fmt.Println(p, SpiralOrder(p.one))
	}
	// output:
	// {[[3] [2]]} [3 2]
	// {[[2 3]]} [2 3]
	// {[[1]]} [1]
	// {[[1 2 3] [4 5 6] [7 8 9]]} [1 2 3 6 9 8 7 4 5]
	// {[[1 2 3 4] [5 6 7 8] [9 10 11 12]]} [1 2 3 4 8 12 11 10 9 5 6 7]
}
