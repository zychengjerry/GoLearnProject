package main

import (
	"fmt"
)

func main() {
	// %d 表示整型数字，%s 表示字符串
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

	var a string = "Jerry Cheng"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 没有初始化就是0值
	var d int
	fmt.Println(d)

	// boolean零值为 false
	var e bool
	fmt.Println(e)

	var i int
	var f float64
	var n bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, n, s)

}
