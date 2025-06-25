package pgsql

import (
	"database/sql"
	"example/clean-arch/internal/entity"
	"example/clean-arch/internal/repository"
	"log"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepository creates a new instance of userRepo
func NewUserRepo(db *sqlx.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) All() ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.Select(&users, "SELECT * FROM users")

	log.Println(err)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepo) GetById(id string) (*entity.User, error) {
	var user entity.User

	err := r.db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) GetByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Create(user *entity.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id",
		user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
