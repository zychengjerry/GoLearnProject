package main

import "fmt"

// 在变量 a 等于 15 的时候跳过本次循环执行下一次循环
func main() {
	/* 定义局部变量 */
	var a int = 10

	/* for 循环 */
	for a < 20 {
		if a == 15 {
			/* 跳过此次循环 */
			a = a + 1
			continue
		}
		fmt.Printf("a 的值为 : %d\n", a)
		a++
	}
}
