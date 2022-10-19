package search

import (
	"fmt"
)

type question74 struct {
	para74
	ans74
}

// para 是参数
// one 代表第一个参数
type para74 struct {
	matrix [][]int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans74 struct {
	one bool
}

func ExampleSearchMatrix() {

	qs := []question74{

		{
			para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 3},
			ans74{true},
		},

		{
			para74{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}, 13},
			ans74{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans74, q.para74
		fmt.Println(p, SearchMatrix(p.matrix, p.target))
	}
	// output:
	// {[[1 3 5 7] [10 11 16 20] [23 30 34 50]] 3} true
	// {[[1 3 5 7] [10 11 16 20] [23 30 34 50]] 13} false
}
