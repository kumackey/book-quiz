package usecase

import "github.com/book-quiz/entity"

type IdGenerator interface {
	Generate() entity.Id
}
