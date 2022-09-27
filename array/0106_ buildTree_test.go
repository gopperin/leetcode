package array

import (
	"fmt"

	"github.com/gopperin/leetcode/structures"
)

type question106 struct {
	para106
	ans106
}

// para 是参数
// one 代表第一个参数
type para106 struct {
	inorder   []int
	postorder []int
}

// ans 是答案
// one 代表第一个答案
type ans106 struct {
	one []int
}

func ExampleBuildTree01() {

	qs := []question106{

		{
			para106{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			ans106{[]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}},
		},
	}

	for _, q := range qs {
		_, p := q.ans106, q.para106
		fmt.Println(p)
		fmt.Println(structures.Tree2ints(BuildTree01(p.inorder, p.postorder)))
	}

	// Output:
	// {[9 3 15 20 7] [9 15 7 20 3]}
	// [3 9 15 -9223372036854775808 -9223372036854775808 -9223372036854775808 20 7]
}
