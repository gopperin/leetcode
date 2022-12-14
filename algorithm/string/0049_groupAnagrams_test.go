package string

import (
	"fmt"
)

type question49 struct {
	para49
	ans49
}

// para 是参数
// one 代表第一个参数
type para49 struct {
	one []string
}

// ans 是答案
// one 代表第一个答案
type ans49 struct {
	one [][]string
}

func ExampleGroupAnagrams() {

	qs := []question49{

		{
			para49{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			ans49{[][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		},
	}

	for _, q := range qs {
		_, p := q.ans49, q.para49
		fmt.Println(p, GroupAnagrams(p.one))
	}
	// {[eat tea tan ate nat bat]} [[eat tea ate] [tan nat] [bat]]
}
