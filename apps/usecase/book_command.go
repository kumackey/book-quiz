package usecase

import (
	"context"
	"github.com/bookuiz-apps/entity"
)

type BookCommand interface {
	exec() (entity.BookId, error)
}

type bookCommand struct {
	bookRepo BookRepository
	uuidGen  UuidGenerator
}

func (c *bookCommand) exec(ctx context.Context) (entity.BookId, error) {
	bookId := entity.BookId(c.uuidGen.New())
	bookRegistered := entity.NewBookRegistered(bookId)

	err := c.bookRepo.Persist(ctx, bookRegistered)
	if err != nil {
		return "", err
	}

	return bookId, nil
}
