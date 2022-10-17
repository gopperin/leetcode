package list

import (
	"fmt"

	"github.com/gopperin/leetcode/structures"
)

type question21 struct {
	para21
	ans21
}

// para 是参数
// one 代表第一个参数
type para21 struct {
	one     []int
	another []int
}

// ans 是答案
// one 代表第一个答案
type ans21 struct {
	one []int
}

func ExampleMergeTwoSortedLists() {

	qs := []question21{

		{
			para21{[]int{}, []int{}},
			ans21{[]int{}},
		},

		{
			para21{[]int{1}, []int{1}},
			ans21{[]int{1, 1}},
		},

		{
			para21{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			para21{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			ans21{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			para21{[]int{1}, []int{9, 9, 9, 9, 9}},
			ans21{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para21{[]int{9, 9, 9, 9, 9}, []int{1}},
			ans21{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para21{[]int{2, 3, 4}, []int{4, 5, 6}},
			ans21{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			para21{[]int{1, 3, 8}, []int{1, 7}},
			ans21{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	for _, q := range qs {
		_, p := q.ans21, q.para21
		fmt.Println(p, structures.List2Ints(MergeTwoSortedLists(structures.Ints2List(p.one), structures.Ints2List(p.another))))
	}

	// Output:
	// {[] []} []
	// {[1] [1]} [1 1]
	// {[1 2 3 4] [1 2 3 4]} [1 1 2 2 3 3 4 4]
	// {[1 2 3 4 5] [1 2 3 4 5]} [1 1 2 2 3 3 4 4 5 5]
	// {[1] [9 9 9 9 9]} [1 9 9 9 9 9]
	// {[9 9 9 9 9] [1]} [1 9 9 9 9 9]
	// {[2 3 4] [4 5 6]} [2 3 4 4 5 6]
	// {[1 3 8] [1 7]} [1 1 3 7 8]
}
