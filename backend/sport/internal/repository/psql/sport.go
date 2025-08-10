package psql

import (
	"context"
	"database/sql"

	//"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	"github.com/DexScen/SuSuSport/backend/sport/internal/errors"
	_ "github.com/lib/pq"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u *Users) GetPassword(ctx context.Context, login string) (string, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	statement, err := tr.Prepare("SELECT password FROM users WHERE login=$1")
	if err != nil {
		tr.Rollback()
		return "", err
	}
	defer statement.Close()

	var password string
	err = statement.QueryRow(login).Scan(&password)
	if err != nil {
		tr.Rollback()
		if err == sql.ErrNoRows {
			return "", errors.ErrUserNotFound
		}
		return "", err
	}

	if err := tr.Commit(); err != nil {
		return "", err
	}

	return password, nil
}

func (u *Users) GetRole(ctx context.Context, login string) (string, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return "", err
	}
	statement, err := tr.Prepare("SELECT role FROM users WHERE login=$1")
	if err != nil {
		tr.Rollback()
		return "", err
	}
	defer statement.Close()

	var role string
	err = statement.QueryRow(login).Scan(&role)
	if err != nil {
		tr.Rollback()
		if err == sql.ErrNoRows {
			return "", errors.ErrUserNotFound
		}
		return "", err
	}

	if err := tr.Commit(); err != nil {
		return "", err
	}

	return role, nil
}