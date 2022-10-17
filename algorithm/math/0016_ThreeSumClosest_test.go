package math

import (
	"fmt"
)

type question16 struct {
	para16
	ans16
}

// para 是参数
// one 代表第一个参数
type para16 struct {
	a      []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans16 struct {
	one int
}

func ExampleThreeSumClosest() {

	qs := []question16{

		{
			para16{[]int{-1, 0, 1, 1, 55}, 3},
			ans16{2},
		},

		{
			para16{[]int{0, 0, 0}, 1},
			ans16{0},
		},

		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},

		{
			para16{[]int{1, 1, -1}, 0},
			ans16{1},
		},

		{
			para16{[]int{-1, 2, 1, -4}, 1},
			ans16{2},
		},
	}

	for _, q := range qs {
		_, p := q.ans16, q.para16
		fmt.Println(p, ThreeSumClosest(p.a, p.target))
	}
	// Output:
	// {[-1 0 1 1 55] 3} 2
	// {[0 0 0] 1} 0
	// {[-4 -1 1 2] 1} 2
	// {[-1 1 1] 0} 1
	// {[-4 -1 1 2] 1} 2
}
