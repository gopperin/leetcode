package string

import (
	"fmt"
)

type question32 struct {
	para32
	ans32
}

// para 是参数
// one 代表第一个参数
type para32 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans32 struct {
	one int
}

func ExampleLongestValidParentheses() {

	qs := []question32{

		{
			para32{"(()"},
			ans32{2},
		},

		{
			para32{")()())"},
			ans32{4},
		},

		{
			para32{"()(())"},
			ans32{6},
		},

		{
			para32{"()(())))"},
			ans32{6},
		},

		{
			para32{"()(()"},
			ans32{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans32, q.para32
		fmt.Println(p, LongestValidParentheses(p.s))
	}
	// output:
	// {(()} 2
	// {)()())} 4
	// {()(())} 6
	// {()(())))} 6
	// {()(()} 2
}
