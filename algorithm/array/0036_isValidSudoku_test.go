package array

import (
	"fmt"
)

type question36 struct {
	para36
	ans36
}

// para 是参数
// one 代表第一个参数
type para36 struct {
	s [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans36 struct {
	s bool
}

func ExampleIsValidSudoku1() {

	qs := []question36{

		{
			para36{[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			ans36{true},
		},

		{
			para36{[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			ans36{false},
		},

		{
			para36{[][]byte{
				{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
				{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'9', '.', '.', '.', '.', '.', '.', '.', '.'}}},
			ans36{true},
		},
	}

	for _, q := range qs {
		_, p := q.ans36, q.para36
		fmt.Println(p, IsValidSudoku1(p.s))
	}
	// output:
	// {[[53 51 46 46 55 46 46 46 46] [54 46 46 49 57 53 46 46 46] [46 57 56 46 46 46 46 54 46] [56 46 46 46 54 46 46 46 51] [52 46 46 56 46 51 46 46 49] [55 46 46 46 50 46 46 46 54] [46 54 46 46 46 46 50 56 46] [46 46 46 52 49 57 46 46 53] [46 46 46 46 56 46 46 55 57]]} true
	// {[[56 51 46 46 55 46 46 46 46] [54 46 46 49 57 53 46 46 46] [46 57 56 46 46 46 46 54 46] [56 46 46 46 54 46 46 46 51] [52 46 46 56 46 51 46 46 49] [55 46 46 46 50 46 46 46 54] [46 54 46 46 46 46 50 56 46] [46 46 46 52 49 57 46 46 53] [46 46 46 46 56 46 46 55 57]]} false
	// {[[46 56 55 54 53 52 51 50 49] [50 46 46 46 46 46 46 46 46] [51 46 46 46 46 46 46 46 46] [52 46 46 46 46 46 46 46 46] [53 46 46 46 46 46 46 46 46] [54 46 46 46 46 46 46 46 46] [55 46 46 46 46 46 46 46 46] [56 46 46 46 46 46 46 46 46] [57 46 46 46 46 46 46 46 46]]} true
}
