package math

import (
	"fmt"
)

type question63 struct {
	para63
	ans63
}

// para 是参数
// one 代表第一个参数
type para63 struct {
	og [][]int
}

// ans 是答案
// one 代表第一个答案
type ans63 struct {
	one int
}

func ExampleUniquePathsWithObstacles() {

	qs := []question63{

		{
			para63{[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			}},
			ans63{2},
		},

		{
			para63{[][]int{
				{0, 0},
				{1, 1},
				{0, 0},
			}},
			ans63{0},
		},

		{
			para63{[][]int{
				{0, 1, 0, 0},
				{1, 0, 0, 0},
				{0, 0, 0, 0},
			}},
			ans63{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans63, q.para63
		fmt.Println(p, UniquePathsWithObstacles(p.og))
	}
	// output:
	// {[[0 0 0] [0 1 0] [0 0 0]]} 2
	// {[[0 0] [1 1] [0 0]]} 0
	// {[[0 1 0 0] [1 0 0 0] [0 0 0 0]]} 0
}
