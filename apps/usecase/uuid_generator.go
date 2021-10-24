package usecase

type UuidGenerator interface {
	New() Uuid
}

type Uuid string
