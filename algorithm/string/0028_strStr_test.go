package string

import (
	"fmt"
)

type question28 struct {
	para28
	ans28
}

// para 是参数
// one 代表第一个参数
type para28 struct {
	s string
	p string
}

// ans 是答案
// one 代表第一个答案
type ans28 struct {
	one int
}

func ExampleStrStr() {

	qs := []question28{

		{
			para28{"abab", "ab"},
			ans28{0},
		},

		{
			para28{"hello", "ll"},
			ans28{2},
		},

		{
			para28{"", "abc"},
			ans28{0},
		},

		{
			para28{"abacbabc", "abc"},
			ans28{5},
		},

		{
			para28{"abacbabc", "abcd"},
			ans28{-1},
		},

		{
			para28{"abacbabc", ""},
			ans28{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans28, q.para28
		fmt.Println(p, StrStr(p.s, p.p))
	}
	// output:
	// {abab ab} 0
	// {hello ll} 2
	// { abc} -1
	// {abacbabc abc} 5
	// {abacbabc abcd} -1
	// {abacbabc } 0
}
