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
}
