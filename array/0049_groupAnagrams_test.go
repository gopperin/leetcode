package array

import (
	"fmt"
)

func ExampleGroupAnagrams() {

	input := [...]string{"eat", "tea", "tan", "ate", "nat", "bat"}

	ret := GroupAnagrams(input[:])

	fmt.Println(ret)

	// Output:
	// [[eat tea ate] [tan nat] [bat]]
}
