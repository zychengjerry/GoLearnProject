package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Hello")
	say("World")
}

/*
World
Hello
Hello
World
World
Hello
Hello
World
World
*/

/* 输出的 hello 和 world 是没有固定先后顺序。
因为它们是两个 goroutine 在执行 */
