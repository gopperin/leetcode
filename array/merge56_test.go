package array

import (
	"fmt"
)

func ExampleMerge56() {

	// Step 1: 创建数组
	values := []Interval{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := Interval{1, 3}
	row2 := Interval{2, 6}
	row3 := Interval{8, 10}
	row4 := Interval{15, 18}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	ret := Merge56(values)

	fmt.Println(ret)

	// Output:
	// [{1 6} {8 10} {15 18}]
}
