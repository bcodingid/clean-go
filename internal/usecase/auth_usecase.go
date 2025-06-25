package usecase

import (
	"errors"
	"example/clean-arch/internal/entity"
	"example/clean-arch/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// define the interfase of usecase
type AuthUsecase interface {
	Login(params entity.LoginParams) (string, error)
	Register(entity.RegisterParams) error
}

// define a user usecase struct
type authUsecase struct {
	repo      repository.UserRepository
	jwtSecret string
}

// struct (class) instantiation on golang
func NewAuthUsecase(repo repository.UserRepository, secret string) AuthUsecase {
	return &authUsecase{
		repo:      repo,
		jwtSecret: secret,
	}
}

// Login method to authenticate user and return JWT token
func (uc *authUsecase) Login(params entity.LoginParams) (string, error) {
	user, err := uc.repo.GetByEmail(params.Email)

	if err != nil {
		return "", errors.New("email not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))

	if err != nil {
		return "", errors.New("invalid password")
	}

	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, _ := token.SignedString([]byte(uc.jwtSecret))

	return signed, nil
}

// Register method to create a new user
func (uc *authUsecase) Register(params entity.RegisterParams) error {
	_, err := uc.repo.GetByEmail(params.Email)

	if err != nil {
		return errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	user := &entity.User{
		Username: params.Username,
		Email:    params.Email,
		Password: string(hashedPassword),
	}

	err = uc.repo.Create(user)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
