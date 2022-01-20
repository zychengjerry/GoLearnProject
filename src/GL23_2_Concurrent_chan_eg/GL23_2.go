package main

import "fmt"

/*
通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 7,2,8
	go sum(s[len(s)/2:], c) // -9,4,0
	x, y := <-c, <-c        // 从通道 c 中接收

	fmt.Println(x, y, x+y)
}

/*
-5 17 12
*/
