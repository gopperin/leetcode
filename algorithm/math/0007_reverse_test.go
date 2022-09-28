package math

import (
	"fmt"
)

type question7 struct {
	para7
	ans7
}

// para 是参数
// one 代表第一个参数
type para7 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans7 struct {
	one int
}

func ExampleReverse7() {

	qs := []question7{

		{
			para7{321},
			ans7{123},
		},

		{
			para7{-123},
			ans7{-321},
		},

		{
			para7{120},
			ans7{21},
		},

		{
			para7{1534236469},
			ans7{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans7, q.para7
		fmt.Println(p.one, Reverse7(p.one))
	}

	// Output:
	// 	321 123
	// -123 -321
	// 120 21
	// 1534236469 0
}
