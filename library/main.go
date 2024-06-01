package main

import "fmt"

type Item interface {
	borrow() string
	returnItem() string
}

type Book struct {
	title  string
	author string
}

type DVD struct {
	title    string
	director string
}

func (b Book) borrow() string {
	return "Book borrowed successfully!"
}

func (b Book) returnItem() string {
	return "Book returned successfully!"
}

func (d DVD) borrow() string {
	return "DVD borrowed successfully!"
}

func (d DVD) returnItem() string {
	return "DVD returned successfully!"
}

func main() {
	book := Book{title: "The Hobbit", author: "J.R.R. Tolkien"}
	dvd := DVD{title: "Inception", director: "Christopher Nolan"}

	fmt.Println(book.borrow())
	fmt.Println(dvd.borrow())

	fmt.Println(book.returnItem())
	fmt.Println(dvd.returnItem())
}
