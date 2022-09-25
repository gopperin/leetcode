package array

import (
	"fmt"
)

func ExampleUniquePathsWithObstacles() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{0, 0, 0}
	row2 := []int{0, 1, 0}
	row3 := []int{0, 0, 0}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	ret := UniquePathsWithObstacles(values)

	fmt.Println(ret)

	// Output:
	// 2
}
