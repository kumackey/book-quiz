package entity

type BookRegistered struct {
	bookId BookId
}

func NewBookRegistered(id BookId) BookRegistered {
	return BookRegistered{bookId: id}
}
