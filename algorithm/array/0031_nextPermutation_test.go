package array

import (
	"fmt"
)

type question31 struct {
	para31
	ans31
}

// para 是参数
// one 代表第一个参数
type para31 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans31 struct {
	one []int
}

func ExampleNextPermutation() {

	qs := []question31{

		{
			para31{[]int{1, 2, 3}},
			ans31{[]int{1, 3, 2}},
		},

		{
			para31{[]int{3, 2, 1}},
			ans31{[]int{1, 2, 3}},
		},

		{
			para31{[]int{1, 1, 5}},
			ans31{[]int{1, 5, 1}},
		},

		{
			para31{[]int{1}},
			ans31{[]int{1}},
		},
	}

	for _, q := range qs {
		_, p := q.ans31, q.para31
		NextPermutation(p.nums)
		fmt.Println(p.nums)
	}

	// output:
	// [1 3 2]
	// [1 2 3]
	// [1 5 1]
	// [1]
}
