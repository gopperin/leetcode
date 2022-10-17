package array

import (
	"fmt"
)

type question26 struct {
	para26
	ans26
}

// para 是参数
// one 代表第一个参数
type para26 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans26 struct {
	one int
}

func ExampleRemoveDuplicates() {

	qs := []question26{

		{
			para26{[]int{1, 1, 2}},
			ans26{2},
		},

		{
			para26{[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
			ans26{5},
		},

		{
			para26{[]int{0, 0, 0, 0, 0}},
			ans26{1},
		},

		{
			para26{[]int{1}},
			ans26{1},
		},
	}

	for _, q := range qs {
		_, p := q.ans26, q.para26
		fmt.Println(p.one, RemoveDuplicates(p.one))
	}
	// output:
	// [1 2 2] 2
	// [0 1 2 3 4 1 2 3 4 4] 5
	// [0 0 0 0 0] 1
	// [1] 1
}
