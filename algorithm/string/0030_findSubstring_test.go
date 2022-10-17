package string

import (
	"fmt"
)

type question30 struct {
	para30
	ans30
}

// para 是参数
// one 代表第一个参数
type para30 struct {
	one string
	two []string
}

// ans 是答案
// one 代表第一个答案
type ans30 struct {
	one []int
}

func ExampleFindSubstring() {

	qs := []question30{

		{
			para30{"aaaaaaaa", []string{"aa", "aa", "aa"}},
			ans30{[]int{0, 1, 2}},
		},

		{
			para30{"barfoothefoobarman", []string{"foo", "bar"}},
			ans30{[]int{0, 9}},
		},

		{
			para30{"wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}},
			ans30{[]int{}},
		},

		{
			para30{"goodgoodgoodgoodgood", []string{"good"}},
			ans30{[]int{0, 4, 8, 12, 16}},
		},

		{
			para30{"barofoothefoolbarman", []string{"foo", "bar"}},
			ans30{[]int{}},
		},

		{
			para30{"bbarffoothefoobarman", []string{"foo", "bar"}},
			ans30{[]int{}},
		},

		{
			para30{"ooroodoofoodtoo", []string{"foo", "doo", "roo", "tee", "oo"}},
			ans30{[]int{}},
		},

		{
			para30{"abc", []string{"a", "b", "c"}},
			ans30{[]int{0}},
		},

		{
			para30{"a", []string{"b"}},
			ans30{[]int{}},
		},

		{
			para30{"ab", []string{"ba"}},
			ans30{[]int{}},
		},

		{
			para30{"n", []string{}},
			ans30{[]int{}},
		},
	}

	for _, q := range qs {
		_, p := q.ans30, q.para30
		fmt.Println(p, FindSubstring(p.one, p.two))
	}
	// output:
	// {aaaaaaaa [aa aa aa]} [0 1 2]
	// {barfoothefoobarman [foo bar]} [0 9]
	// {wordgoodgoodgoodbestword [word good best word]} []
	// {goodgoodgoodgoodgood [good]} [0 4 8 12 16]
	// {barofoothefoolbarman [foo bar]} []
	// {bbarffoothefoobarman [foo bar]} [11]
	// {ooroodoofoodtoo [foo doo roo tee oo]} []
	// {abc [a b c]} [0]
	// {a [b]} []
	// {ab [ba]} []
	// {n []} []
}
