package array

import (
	"fmt"
)

type question46 struct {
	para46
	ans46
}

// para 是参数
// one 代表第一个参数
type para46 struct {
	s []int
}

// ans 是答案
// one 代表第一个答案
type ans46 struct {
	one [][]int
}

func ExamplePermute() {

	qs := []question46{

		{
			para46{[]int{1, 2, 3}},
			ans46{[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},
	}

	for _, q := range qs {
		_, p := q.ans46, q.para46
		fmt.Println(p, Permute(p.s))
	}
	// output:
	// {[1 2 3]} [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]
}
