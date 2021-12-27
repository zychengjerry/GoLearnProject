package main

import "fmt"

func main() {
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 appped() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
}

/*
Row 1
[1 2 3]
Row 2
[4 5 6]
第一个元素为：
1
*/
