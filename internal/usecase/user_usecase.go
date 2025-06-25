package usecase

import (
	"example/clean-arch/internal/entity"
	"example/clean-arch/internal/repository"
)

type UserUsecase interface {
	All() ([]*entity.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *userUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (uc *userUsecase) All() ([]*entity.User, error) {
	users, err := uc.repo.All()

	if err != nil {
		return nil, err
	}
	return users, nil
}
