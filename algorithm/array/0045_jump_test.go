package array

import (
	"fmt"
)

type question45 struct {
	para45
	ans45
}

// para 是参数
// one 代表第一个参数
type para45 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans45 struct {
	one int
}

func ExampleJump() {

	qs := []question45{

		{
			para45{[]int{2, 3, 1, 1, 4}},
			ans45{2},
		},

		{
			para45{[]int{2, 3, 0, 1, 4}},
			ans45{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans45, q.para45
		fmt.Println(p, Jump(p.nums))
	}
	// output:
	// {[2 3 1 1 4]} 2
	// {[2 3 0 1 4]} 2
}
