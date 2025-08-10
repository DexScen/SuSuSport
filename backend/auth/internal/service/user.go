package service

import (
	"context"
	"errors"

	//"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	e "github.com/DexScen/SuSuSport/backend/auth/internal/errors"
)

type UsersRepository interface {
	GetPassword(ctx context.Context, login string) (string, error)
	GetUser(ctx context.Context, login string) (*domain.User, error)
}

type Users struct {
	repo UsersRepository
}

func NewUsers(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

func (u *Users) LogIn(ctx context.Context, login, password string) (*domain.User, error) {
	passwordFromDB, err := u.repo.GetPassword(ctx, login)

	if err != nil {
		if errors.Is(err, e.ErrUserNotFound) {
			return nil, e.ErrUserNotFound
		}
		return nil, err
	}

	if password == passwordFromDB{
		user, err := u.repo.GetUser(ctx, login)
		if err != nil{
			return nil, err
		}
		return user, nil
	}
	return nil, e.ErrWrongPassword
}
