package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Book

	book1.title = "Go language"
	book1.author = "www.runoob.com"
	book1.subject = "Go language book"
	book1.book_id = 6495407

	printBook(&book1)

}

func printBook(book *Book) {
	fmt.Printf("Book title: %s\n", book.title)
	fmt.Printf("Book author: %s\n", book.author)
	fmt.Printf("Book subject: %s\n", book.subject)
	fmt.Printf("Book book_id: %d\n", book.book_id)
}
