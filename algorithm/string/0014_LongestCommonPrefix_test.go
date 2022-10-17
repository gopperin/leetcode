package string

import (
	"fmt"
)

type question14 struct {
	para14
	ans14
}

// para 是参数
type para14 struct {
	strs []string
}

// ans 是答案
type ans14 struct {
	ans string
}

func ExampleLongestCommonPrefix() {

	qs := []question14{

		{
			para14{[]string{"flower", "flow", "flight"}},
			ans14{"fl"},
		},

		{
			para14{[]string{"dog", "racecar", "car"}},
			ans14{""},
		},

		{
			para14{[]string{"ab", "a"}},
			ans14{"a"},
		},
	}

	for _, q := range qs {
		_, p := q.ans14, q.para14
		fmt.Println(p.strs, LongestCommonPrefix(p.strs), "ex")
	}

	// Output:
	// [flower flow flight] fl ex
	// [dog racecar car]  ex
	// [ab a] a ex
}
