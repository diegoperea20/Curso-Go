package book

import (
	"fmt"
)

type Book struct {
	Title string  //si se escribe el atributo en minuscula, es privado
	Author string
	Pages int
}

//para atributos privados se hace get y set
/* func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetTitle(title string) {
	b.title = title
} */


//constructor (simulado)
func NewBook(title string, author string, pages int) *Book {
	return &Book{Title: title, Author: author, Pages: pages}
}

func (b *Book) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}



//herencia

type TextBook struct {
	Book
	editorial string  //si se escribe el atributo en minuscula, es privado
	level string 		//si se escribe el atributo en minuscula, es privado
}
//constructor
func NewTextBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{Book{ title,author,pages}, editorial, level}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n", b.Title, b.Author, b.Pages, b.editorial, b.level)
}