package array

import (
	"fmt"
)

type question51 struct {
	para51
	ans51
}

// para 是参数
// one 代表第一个参数
type para51 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans51 struct {
	one [][]string
}

func ExampleSolveNQueens() {

	qs := []question51{

		{
			para51{4},
			ans51{[][]string{
				{".Q..",
					"...Q",
					"Q...",
					"..Q."},
				{"..Q.",
					"Q...",
					"...Q",
					".Q.."},
			}},
		},
	}

	for _, q := range qs {
		_, p := q.ans51, q.para51
		fmt.Println(p, SolveNQueens(p.one))
	}
	// output:
	// {4} [[.Q.. ...Q Q... ..Q.] [..Q. Q... ...Q .Q..]]
}
