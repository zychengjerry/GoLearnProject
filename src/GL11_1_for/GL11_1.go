package main

import "fmt"

func main() {
	/* 计算1-10的和 */
	sum1 := 0
	for i := 0; i <= 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	// 省略init和post参数
	sum2 := 1
	for sum2 <= 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// 这样写也可以，更像 While 语句形式
	for sum2 <= 10 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// For-each range 循环
	strings := []string{"google", "youtube"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// 无限循环
	sum := 0
	for {
		sum++ // 无限循环下去
	}
	fmt.Println(sum) // 无法输出

	// 要停止无限循环，可以在命令窗口按下ctrl-c

}
