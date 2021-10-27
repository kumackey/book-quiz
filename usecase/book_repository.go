package usecase

import (
	"context"
	"github.com/bookuiz-apps/entity"
)

type BookRepository interface {
	Persist(ctx context.Context, bookRegistered entity.BookRegistered) error
}
