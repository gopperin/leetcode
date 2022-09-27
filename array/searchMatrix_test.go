package array

import (
	"fmt"
)

func ExampleSearchMatrix() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 3, 5, 7}
	row2 := []int{10, 13, 16, 20}
	row3 := []int{23, 30, 34, 50}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	target := 13

	ret := SearchMatrix(values, target)

	fmt.Println(ret)

	// Output:
	// true
}
