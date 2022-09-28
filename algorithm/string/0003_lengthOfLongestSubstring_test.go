package string

import (
	"fmt"
)

type question3 struct {
	para3
	ans3
}

// para 是参数
// one 代表第一个参数
type para3 struct {
	s string
}

// ans 是答案
// one 代表第一个答案
type ans3 struct {
	one int
}

func ExampleLengthOfLongestSubstring() {

	qs := []question3{

		{
			para3{"abcabcbb"},
			ans3{3},
		},

		{
			para3{"bbbbb"},
			ans3{1},
		},

		{
			para3{"pwwkew"},
			ans3{3},
		},

		{
			para3{""},
			ans3{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans3, q.para3
		fmt.Println(p, LengthOfLongestSubstring(p.s))
	}

	// Output:
	// 	{abcabcbb} 3
	// {bbbbb} 1
	// {pwwkew} 3
	// {} 0
}
