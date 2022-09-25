package array

import (
	"fmt"
)

func ExampleSetMatrixZeroes() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{0, 1, 2, 0}
	row2 := []int{3, 4, 5, 2}
	row3 := []int{1, 3, 1, 5}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	SetMatrixZeroes(values)

	fmt.Println(values)

	// Output:
	// [[0 0 0 0] [0 4 5 0] [0 3 1 0]]
}
