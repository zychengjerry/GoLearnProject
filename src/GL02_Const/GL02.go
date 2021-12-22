package main

import "fmt"
import "unsafe"

// 枚举
const (
	x = "xyz"
	y = len(x)
	z = unsafe.Sizeof(x)
)

func main() {
	// 常量
	const length int = 10
	const width int = 5
	var area int

	// 多重赋值
	const a, b, c = 1, false, "str"

	area = length * width
	fmt.Printf("Area is: %d", area)
	println()
	println(a, b, c)
	println(x, y, z)
}
