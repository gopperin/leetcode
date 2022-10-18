package array

import (
	"fmt"
)

type question59 struct {
	para59
	ans59
}

// para 是参数
// one 代表第一个参数
type para59 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans59 struct {
	one [][]int
}

func ExampleGenerateMatrix() {

	qs := []question59{

		{
			para59{3},
			ans59{[][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		},

		{
			para59{4},
			ans59{[][]int{{1, 2, 3, 4}, {12, 13, 14, 5}, {11, 16, 15, 6}, {10, 9, 8, 7}}},
		},
	}

	for _, q := range qs {
		_, p := q.ans59, q.para59
		fmt.Println(p, GenerateMatrix(p.one))
	}
	// output:
	// {3} [[1 2 3] [8 9 4] [7 6 5]]
	// {4} [[1 2 3 4] [12 13 14 5] [11 16 15 6] [10 9 8 7]]
}
