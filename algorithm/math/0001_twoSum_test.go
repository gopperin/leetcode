package math

import (
	"fmt"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func ExampleTwoSum() {

	qs := []question1{
		{
			para1{[]int{3, 2, 4}, 6},
			ans1{[]int{1, 2}},
		},

		{
			para1{[]int{3, 2, 4}, 5},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}, 11},
			ans1{[]int{1, 3}},
		},

		{
			para1{[]int{0, 1}, 1},
			ans1{[]int{0, 1}},
		},

		{
			para1{[]int{0, 3}, 5},
			ans1{[]int{}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Println(p, TwoSum(p.nums, p.target))
	}

	// Output:
	// {[3 2 4] 6} [1 2]
	// {[3 2 4] 5} [0 1]
	// {[0 8 7 3 3 4 2] 11} [1 3]
	// {[0 1] 1} [0 1]
	// {[0 3] 5} []
}
