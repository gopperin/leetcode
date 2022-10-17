package string

import (
	"fmt"
)

type question17 struct {
	para17
	ans17
}

// para 是参数
// one 代表第一个参数
type para17 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans17 struct {
	one []string
}

func ExampleLetterCombinations() {

	qs := []question17{

		{
			para17{"23"},
			ans17{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
	}

	for _, q := range qs {
		_, p := q.ans17, q.para17
		fmt.Println(p, LetterCombinations(p.s))
	}

	// Output:
	// {23} [ad ae af bd be bf cd ce cf]
}
