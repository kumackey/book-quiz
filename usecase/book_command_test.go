package usecase

import (
	"context"
	"github.com/bookuiz-apps/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_Execメソッドは本を作成できる(t *testing.T) {
	id := "123456"

	mockBookRepo := new(mockBookRepository)
	mockBookRepo.On("Persist", mock.Anything, mock.Anything).Return(nil)

	mockIdGen := new(mockIdGenerator)
	mockIdGen.On("Generate").Return(id)

	bookCommand := bookCommand{mockBookRepo, mockIdGen}
	bookId, err := bookCommand.exec(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, id, string(bookId))
}

type mockBookRepository struct {
	mock.Mock
}

func (m *mockBookRepository) Persist(ctx context.Context, bookRegistered entity.BookRegistered) error {
	args := m.Called(ctx, bookRegistered)

	return args.Error(0)
}
