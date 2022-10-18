package backtracking

import (
	"fmt"
)

type question60 struct {
	para60
	ans60
}

// para 是参数
// one 代表第一个参数
type para60 struct {
	n int
	k int
}

// ans 是答案
// one 代表第一个答案
type ans60 struct {
	one string
}

func ExampleGetPermutation() {

	qs := []question60{

		{
			para60{3, 3},
			ans60{"213"},
		},

		{
			para60{4, 9},
			ans60{"2314"},
		},
	}

	for _, q := range qs {
		_, p := q.ans60, q.para60
		fmt.Println(p, GetPermutation(p.n, p.k))
	}
	// output:
	// {3 3} 213
	// {4 9} 2314
}
