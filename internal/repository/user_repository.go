package repository

import "example/clean-arch/internal/entity"

type UserRepository interface {
	All() ([]*entity.User, error)
	GetById(string) (*entity.User, error)
	GetByEmail(string) (*entity.User, error)
	Create(*entity.User) error
}
