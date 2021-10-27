package usecase

import (
	"context"
	"github.com/book-quiz/entity"
)

type BookRepository interface {
	Persist(ctx context.Context, bookRegistered entity.BookRegistered) error
}
