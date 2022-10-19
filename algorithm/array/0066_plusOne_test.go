package array

import (
	"fmt"
)

type question66 struct {
	para66
	ans66
}

// para 是参数
// one 代表第一个参数
type para66 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans66 struct {
	one []int
}

func ExamplePlusOne() {

	qs := []question66{

		{
			para66{[]int{1, 2, 3}},
			ans66{[]int{1, 2, 4}},
		},

		{
			para66{[]int{4, 3, 2, 1}},
			ans66{[]int{4, 3, 2, 2}},
		},

		{
			para66{[]int{9, 9}},
			ans66{[]int{1, 0, 0}},
		},

		{
			para66{[]int{0}},
			ans66{[]int{0}},
		},
	}

	for _, q := range qs {
		_, p := q.ans66, q.para66
		fmt.Println(p, PlusOne(p.one))
	}
	// output:
	// {[1 2 4]} [1 2 4]
	// {[4 3 2 2]} [4 3 2 2]
	// {[1 0]} [1 0 0]
	// {[1]} [1]
}
