package usecase

import (
	"context"
	"github.com/book-quiz/entity"
)

type BookCommand interface {
	exec() (entity.BookId, error)
}

type bookCommand struct {
	bookRepo BookRepository
	idGen    IdGenerator
}

func (c *bookCommand) exec(ctx context.Context) (entity.BookId, error) {
	bookId := entity.BookId(c.idGen.Generate())
	bookRegistered := entity.NewBookRegistered(bookId)

	err := c.bookRepo.Persist(ctx, bookRegistered)
	if err != nil {
		return "", err
	}

	return bookId, nil
}
