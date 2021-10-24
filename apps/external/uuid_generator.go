package external

import (
	"github.com/bookuiz-apps/usecase"
	"github.com/google/uuid"
)

type UuidGenerator struct{}

func New() usecase.Uuid {
	return usecase.Uuid(uuid.New().String())
}
