package array

import (
	"fmt"
)

type question6 struct {
	para6
	ans6
}

// para 是参数
// one 代表第一个参数
type para6 struct {
	s       string
	numRows int
}

// ans 是答案
// one 代表第一个答案
type ans6 struct {
	one string
}

func ExampleZigzag() {

	qs := []question6{

		{
			para6{"PAYPALISHIRING", 3},
			ans6{"PAHNAPLSIIGYIR"},
		},

		{
			para6{"PAYPALISHIRING", 4},
			ans6{"PINALSIGYAHRPI"},
		},

		{
			para6{"A", 1},
			ans6{"A"},
		},
	}

	for _, q := range qs {
		_, p := q.ans6, q.para6
		fmt.Println(p, Zigzag(p.s, p.numRows))
	}
	// Output:
	// 	{PAYPALISHIRING 3} PAHNAPLSIIGYIR
	// {PAYPALISHIRING 4} PINALSIGYAHRPI
	// {A 1} A
}
