package main

import "fmt"

type books struct {
	title  string
	id     int
	author string
	price  float32
}

func printBook(book books) {
	fmt.Println("Title: ", book.title)
	fmt.Println("Author: ", book.author)
	fmt.Println("Price: ", book.price)
	fmt.Println("Id: ", book.id)
}

func main() {
	var book1 books
	var book2 books

	book1.author = "Sam don"
	book1.title = "journey of an Unknown boy"
	book1.id = 12345
	book1.price = 501.01

	book2.author = "Sandy don"
	book2.title = "journey of an known girl"
	book2.id = 12565
	book2.price = 401.01

	fmt.Println("Book1:- ")
	printBook(book1)

	fmt.Println("Book2:- ")
	printBook(book2)
}
