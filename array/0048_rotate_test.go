package array

import (
	"fmt"
	"testing"
)

func ExampleRotate() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{5, 1, 9, 11}
	row2 := []int{2, 4, 8, 10}
	row3 := []int{13, 3, 6, 7}
	row4 := []int{15, 14, 12, 16}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	Rotate(values)

	fmt.Println(values)

	// Output:
	// [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]
}

func BenchmarkRotate(b *testing.B) {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{5, 1, 9, 11}
	row2 := []int{2, 4, 8, 10}
	row3 := []int{13, 3, 6, 7}
	row4 := []int{15, 14, 12, 16}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	for i := 0; i < b.N; i++ {
		Rotate(values)
	}

	// 23014171	        50.05 ns/op	       0 B/op	       0 allocs/op
}

func ExampleRotate1() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{5, 1, 9, 11}
	row2 := []int{2, 4, 8, 10}
	row3 := []int{13, 3, 6, 7}
	row4 := []int{15, 14, 12, 16}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	Rotate1(values)

	fmt.Println(values)

	// Output:
	// [[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]
}

func BenchmarkRotate1(b *testing.B) {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{5, 1, 9, 11}
	row2 := []int{2, 4, 8, 10}
	row3 := []int{13, 3, 6, 7}
	row4 := []int{15, 14, 12, 16}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)
	values = append(values, row4)

	for i := 0; i < b.N; i++ {
		Rotate1(values)
	}

	// 26459953	        44.44 ns/op	       0 B/op	       0 allocs/op
}
