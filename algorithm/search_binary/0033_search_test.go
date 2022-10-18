package search

import (
	"fmt"
)

type question33 struct {
	para33
	ans33
}

// para 是参数
// one 代表第一个参数
type para33 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans33 struct {
	one int
}

func ExampleSearch33() {

	qs := []question33{

		{
			para33{[]int{3, 1}, 1},
			ans33{1},
		},

		{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans33{4},
		},

		{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans33{-1},
		},
	}

	for _, q := range qs {
		_, p := q.ans33, q.para33
		fmt.Println(p, Search33(p.nums, p.target))
	}
	// output:
	// {[3 1] 1} 1
	// {[4 5 6 7 0 1 2] 0} 4
	// {[4 5 6 7 0 1 2] 3} -1
}
