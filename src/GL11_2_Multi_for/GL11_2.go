package main

import "fmt"

/* 以下实例使用循环嵌套来输出 2 到 100 间的素数 */

func main() {
	/* 定义局部变量 */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
}
