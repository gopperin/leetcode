package array

import (
	"fmt"
	"testing"
)

func ExampleMinPathSum() {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 3, 1}
	row2 := []int{1, 5, 1}
	row3 := []int{4, 2, 1}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	ret := MinPathSum(values)

	fmt.Println(ret)

	// Output:
	// 7
}

func BenchmarkMinPathSum(b *testing.B) {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 3, 1}
	row2 := []int{1, 5, 1}
	row3 := []int{4, 2, 1}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	for i := 0; i < b.N; i++ {
		MinPathSum(values)
	}

	// 31945104	        37.72 ns/op	       0 B/op	       0 allocs/op
}

func ExampleMinPathSum1() {
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 3, 1}
	row2 := []int{1, 5, 1}
	row3 := []int{4, 2, 1}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	ret := MinPathSum1(values)

	fmt.Println(ret)

	// Output:
	// 7
}

func BenchmarkMinPathSum1(b *testing.B) {

	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 3, 1}
	row2 := []int{1, 5, 1}
	row3 := []int{4, 2, 1}
	values = append(values, row1)
	values = append(values, row2)
	values = append(values, row3)

	for i := 0; i < b.N; i++ {
		MinPathSum1(values)
	}

	// 3507390	       398.6 ns/op	     152 B/op	       4 allocs/op
}
