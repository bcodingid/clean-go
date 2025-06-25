package usecase

import (
	"example/clean-arch/internal/entity"
	"example/clean-arch/internal/repository"
)

type UserUsecase interface {
	All() ([]*entity.User, error)
	Get(string) (*entity.User, error)
	Create(*entity.User) error
	Update(*entity.User) error
	Delete(string) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
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

func (uc *userUsecase) Get(id string) (*entity.User, error) {
	user, err := uc.repo.GetById(id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUsecase) Create(user *entity.User) error {

	return nil
}

func (uc *userUsecase) Update(user *entity.User) error {
	return nil
}

func (uc *userUsecase) Delete(id string) error {
	return nil
}
