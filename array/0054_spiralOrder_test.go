package array

import (
	"fmt"
	"testing"
)

func ExampleSpiralOrder() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3, 4}
	row2 := []int{5, 6, 7, 8}
	row3 := []int{9, 10, 11, 12}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	ret := SpiralOrder(values)

	fmt.Println(ret)

	// Output:
	// [1 2 3 4 8 12 11 10 9 5 6 7]
}

func BenchmarkSpiralOrder(b *testing.B) {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3, 4}
	row2 := []int{5, 6, 7, 8}
	row3 := []int{9, 10, 11, 12}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	for i := 0; i < b.N; i++ {
		SpiralOrder(values)
	}

	// 1484229	      1476 ns/op	     424 B/op	       9 allocs/op
}

func ExampleSpiralOrder2() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3, 4}
	row2 := []int{5, 6, 7, 8}
	row3 := []int{9, 10, 11, 12}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	ret := SpiralOrder2(values)

	fmt.Println(ret)

	// Output:
	// [1 2 3 4 8 12 11 10 9 5 6 7]
}

func BenchmarkSpiralOrder2(b *testing.B) {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3, 4}
	row2 := []int{5, 6, 7, 8}
	row3 := []int{9, 10, 11, 12}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	for i := 0; i < b.N; i++ {
		SpiralOrder2(values)
	}

	// 3087230	       627.8 ns/op	     248 B/op	       5 allocs/op
}
