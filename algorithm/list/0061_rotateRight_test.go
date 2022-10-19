package list

import (
	"fmt"

	"github.com/gopperin/leetcode/structures"
)

type question61 struct {
	para61
	ans61
}

// para 是参数
// one 代表第一个参数
type para61 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans61 struct {
	one []int
}

func ExampleRotateRight() {

	qs := []question61{

		{
			para61{[]int{1, 2, 3, 4, 5}, 2},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{1, 2, 3, 4, 5}, 3},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{0, 1, 2}, 4},
			ans61{[]int{2, 0, 1}},
		},

		{
			para61{[]int{1, 1, 1, 2}, 3},
			ans61{[]int{1, 1, 2, 1}},
		},

		{
			para61{[]int{1}, 10},
			ans61{[]int{1}},
		},

		{
			para61{[]int{}, 100},
			ans61{[]int{}},
		},
	}

	for _, q := range qs {
		_, p := q.ans61, q.para61
		fmt.Println(p, structures.List2Ints(RotateRight(structures.Ints2List(p.one), p.k)))
	}
	// output:
	// {[1 2 3 4 5] 2} [4 5 1 2 3]
	// {[1 2 3 4 5] 3} [3 4 5 1 2]
	// {[0 1 2] 4} [2 0 1]
	// {[1 1 1 2] 3} [1 1 2 1]
	// {[1] 10} [1]
	// {[] 100} []
}
