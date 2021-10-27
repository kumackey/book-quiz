package adapter

import (
	"github.com/book-quiz/entity"
	"github.com/google/uuid"
)

type IdGenerator struct{}

func (u *IdGenerator) Generate() entity.Id {
	return entity.Id(uuid.NewString())
}
