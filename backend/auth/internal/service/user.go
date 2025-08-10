package service

import (
	"context"
	"errors"

	//"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	e "github.com/DexScen/SuSuSport/backend/auth/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository interface {
	GetPassword(ctx context.Context, login string) (string, error)
	GetRole(ctx context.Context, login string) (string, error)
}

type Users struct {
	repo UsersRepository
}

func NewUsers(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func (u *Users) LogIn(ctx context.Context, login, password string) (string, error) {
	passwordHash, err := u.repo.GetPassword(ctx, login)

	if err != nil {
		if errors.Is(err, e.ErrUserNotFound) {
			return "", e.ErrUserNotFound
		}
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return "", e.ErrWrongPassword
	}
	return u.repo.GetRole(ctx, login)
}
