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
}

type BookRepository interface {
	persist(ctx context.Context, bookRegistered entity.BookRegistered) error
}

func (bc *bookCommand) exec() (entity.BookId, error) {
	return "", nil
}
