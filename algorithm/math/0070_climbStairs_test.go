package math

import (
	"fmt"
)

type question70 struct {
	para70
	ans70
}

// para 是参数
// one 代表第一个参数
type para70 struct {
	n int
}

// ans 是答案
// one 代表第一个答案
type ans70 struct {
	one int
}

func ExampleClimbStairs() {

	qs := []question70{

		{
			para70{2},
			ans70{2},
		},

		{
			para70{3},
			ans70{3},
		},
	}

	for _, q := range qs {
		_, p := q.ans70, q.para70
		fmt.Println(p, ClimbStairs(p.n))
	}
	// output:
	// {2} 2
	// {3} 3
}
