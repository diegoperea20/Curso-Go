package main

import (
	"dependenciasProyecto/POO/book"
)

func main() {

	mybookOtraForma := book.NewBook("The Lord of the Rings", "JRR Tolkien", 700)
	mybookOtraForma.PrintInfo()
	println("---------Otra forma---------")

	var myBook = book.Book{
		Title: "The Lord of the Rings",
		Author: "JRR Tolkien",
		Pages: 700,
	}
	myBook.PrintInfo()

	println("---------Herencia---------")

	myTextBook := book.NewTextBook("The Lord of the Rings 2" , "JRR Tolkien 2", 800, "HarperCollins", "Advanced")

	myTextBook.PrintInfo()
}