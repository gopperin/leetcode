package hash_table

import (
	"fmt"
)

type question73 struct {
	para73
	ans73
}

// para 是参数
// one 代表第一个参数
type para73 struct {
	matrix [][]int
}

// ans 是答案
// one 代表第一个答案
type ans73 struct {
}

func ExampleSetZeroes() {

	qs := []question73{

		{
			para73{[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			}},
			ans73{},
		},
	}

	for _, q := range qs {
		_, p := q.ans73, q.para73
		SetZeroes(p.matrix)
		fmt.Println(p)
	}
	// output:
	// {[[0 0 0 0] [0 4 5 0] [0 3 1 0]]}
}
