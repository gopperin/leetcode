package string

import (
	"fmt"
)

type question12 struct {
	para12
	ans12
}

// para 是参数
// one 代表第一个参数
type para12 struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans12 struct {
	one string
}

func ExampleIntToRoman() {

	qs := []question12{

		{
			para12{3},
			ans12{"III"},
		},

		{
			para12{4},
			ans12{"IV"},
		},

		{
			para12{9},
			ans12{"IX"},
		},

		{
			para12{58},
			ans12{"LVIII"},
		},

		{
			para12{1994},
			ans12{"MCMXCIV"},
		},
		{
			para12{123},
			ans12{"CXXIII"},
		},

		{
			para12{120},
			ans12{"CXX"},
		},
	}

	for _, q := range qs {
		_, p := q.ans12, q.para12
		fmt.Println(p.one, IntToRoman(p.one))
	}

	// Output:
	// 3 III
	// 4 IV
	// 9 IX
	// 58 LVIII
	// 1994 MCMXCIV
	// 123 CXXIII
	// 120 CXX
}
