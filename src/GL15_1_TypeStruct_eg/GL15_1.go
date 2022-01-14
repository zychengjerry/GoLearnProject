package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"How to Learn Go", "Jerry", "Basic", 143204})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "How to Learn Go", author: "Jerry", subject: "Basic", book_id: 143204})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "How to Learn Go", author: "Jerry"})
}

/*
{How to Learn Go Jerry Basic 143204}
{How to Learn Go Jerry Basic 143204}
{How to Learn Go Jerry  0}
*/
