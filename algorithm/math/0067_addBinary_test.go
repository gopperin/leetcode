package math

import (
	"fmt"
)

type question67 struct {
	para67
	ans67
}

// para 是参数
// one 代表第一个参数
type para67 struct {
	a string
	b string
}

// ans 是答案
// one 代表第一个答案
type ans67 struct {
	one string
}

func ExampleAddBinary() {

	qs := []question67{

		{
			para67{"11", "1"},
			ans67{"100"},
		},

		{
			para67{"1010", "1011"},
			ans67{"10101"},
		},
	}

	for _, q := range qs {
		_, p := q.ans67, q.para67
		fmt.Println(p, AddBinary(p.a, p.b))
	}
	// output:
	// {11 1} 100
	// {1010 1011} 10101
}
