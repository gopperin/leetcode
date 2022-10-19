package math

import (
	"fmt"
)

type question69 struct {
	para69
	ans69
}

// para 是参数
// one 代表第一个参数
type para69 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans69 struct {
	one int
}

func ExampleSqrt() {

	qs := []question69{

		{
			para69{4},
			ans69{2},
		},

		{
			para69{8},
			ans69{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans69, q.para69
		fmt.Println(p, Sqrt(p.one))
	}
	// output:
	// {4} 2
	// {8} 2
}
