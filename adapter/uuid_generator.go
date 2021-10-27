package adapter

import (
	"github.com/bookuiz-apps/entity"
	"github.com/google/uuid"
)

type IdGenerator struct{}

func (u *IdGenerator) Generate() entity.Id {
	return entity.Id(uuid.NewString())
}
