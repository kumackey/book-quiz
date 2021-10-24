package usecase

import (
	"github.com/bookuiz-apps/entity"
	"github.com/stretchr/testify/mock"
)

type mockIdGenerator struct {
	mock.Mock
}

func (m *mockIdGenerator) Generate() entity.Id {
	args := m.Called()

	return entity.Id(args.String(0))
}
