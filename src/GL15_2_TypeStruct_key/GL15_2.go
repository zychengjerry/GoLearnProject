package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books

	Book1.title = "How to learn Go"
	Book1.author = "Jerry"
	Book1.subject = "Basic"
	Book1.book_id = 12345

	Book2.title = "How to learn Java"
	Book2.author = "Jerry"
	Book2.subject = "Basic"
	Book2.book_id = 12346

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	/* 打印 Book1 信息 */
	printBookPointer(&Book1)

	/* 打印 Book2 信息 */
	printBookPointer(&Book2)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBookPointer(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

/*
Book 1 title : How to learn Go
Book 1 author : Jerry
Book 1 subject : Basic
Book 1 book_id : 12345
Book 2 title : How to learn Java
Book 2 author : Jerry
Book 2 subject : Basic
Book 2 book_id : 12346
Book 1 title : How to learn Go
Book 1 author : Jerry
Book 1 subject : Basic
Book 1 book_id : 12345
Book 2 title : How to learn Java
Book 2 author : Jerry
Book 2 subject : Basic
Book 2 book_id : 12346
Book 1 title : How to learn Go
Book 1 author : Jerry
Book 1 subject : Basic
Book 1 book_id : 12345
Book 2 title : How to learn Java
Book 2 author : Jerry
Book 2 subject : Basic
Book 2 book_id : 12346
*/
