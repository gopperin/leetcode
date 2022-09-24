package array

import (
	"fmt"
)

func ExampleInsert57() {

	// Step 1: 创建数组
	values := []Interval{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := Interval{1, 2}
	row2 := Interval{3, 5}
	row3 := Interval{6, 7}
	row4 := Interval{8, 10}
	row5 := Interval{12, 16}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)
	values = append(values, row5)

	value := Interval{4, 8}

	ret := Insert57(values, value)

	fmt.Println(ret)

	// Output:
	// [{1 2} {3 10} {12 16}]
}
