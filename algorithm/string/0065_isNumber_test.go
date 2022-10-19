package string

import (
	"fmt"
)

func ExampleIsNumber() {

	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0",
			true,
		},

		{
			"e",
			false,
		},

		{
			".",
			false,
		},

		{
			".1",
			true,
		},
	}

	for _, tc := range tcs {
		fmt.Println(tc, IsNumber(tc.s))
	}
	// output:
	// {0 true} true
	// {e false} false
	// {. false} false
	// {.1 true} true
}
