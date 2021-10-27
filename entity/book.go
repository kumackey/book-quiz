package entity

type Book struct {
	bookId  BookId
	eanCode eanCode
	name    bookName
}

type bookName string

type eanCode string
