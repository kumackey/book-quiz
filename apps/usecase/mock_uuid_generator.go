package usecase

import (
	"github.com/stretchr/testify/mock"
)

type mockUuidGenerator struct {
	mock.Mock
}

func (m *mockUuidGenerator) New() Uuid {
	args := m.Called()

	return Uuid(args.String(0))
}
