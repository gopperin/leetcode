package array

import (
	"fmt"
)

func ExampleExistWord() {

	// Step 1: 创建数组
	values := [][]byte{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []byte{'A', 'B', 'C', 'E'}
	row2 := []byte{'S', 'F', 'C', 'S'}
	row3 := []byte{'A', 'D', 'E', 'E'}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	target := "ABCCED"

	ret := ExistWord(values, target)

	fmt.Println(ret)

	// Output:
	// true
}
