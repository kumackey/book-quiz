package usecase

import (
	"context"
	"github.com/bookuiz-apps/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_Execメソッドは本を作成できる(t *testing.T) {
	mockBookRepo := new(mockBookRepository)
	mockBookRepo.On("persist", mock.Anything).Return(nil)

	bookCommand := bookCommand{mockBookRepo}
	_, err := bookCommand.exec()
	assert.NoError(t, err)
}

type mockBookRepository struct {
	mock.Mock
}

func (m *mockBookRepository) persist(ctx context.Context, bookRegistered entity.BookRegistered) error {
	args := m.Called(ctx, bookRegistered)

	return args.Error(0)
}
