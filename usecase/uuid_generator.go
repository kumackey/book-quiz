package usecase

import "github.com/bookuiz-apps/entity"

type IdGenerator interface {
	Generate() entity.Id
}
