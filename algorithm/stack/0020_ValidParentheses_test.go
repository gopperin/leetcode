package stack

import (
	"fmt"
)

type question20 struct {
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func ExampleValidParentheses() {

	qs := []question20{

		{
			para20{"()[]{}"},
			ans20{true},
		},
		{
			para20{"(]"},
			ans20{false},
		},
		{
			para20{"({[]})"},
			ans20{true},
		},
		{
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			para20{"(())]]"},
			ans20{false},
		},
		{
			para20{""},
			ans20{true},
		},
		{
			para20{"["},
			ans20{false},
		},
	}

	for _, q := range qs {
		_, p := q.ans20, q.para20
		fmt.Println(p, ValidParentheses(p.one))
	}

	// Output:
	// {()[]{}} true
	// {(]} false
	// {({[]})} true
	// {(){[({[]})]}} true
	// {((([[[{{{} false
	// {(())]]} false
	// {} true
	// {[} false
}
