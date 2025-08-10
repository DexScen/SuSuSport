package psql

import (
	"context"
	"database/sql"

	//"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	"github.com/DexScen/SuSuSport/backend/auth/internal/errors"
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

func (u *Users) GetUser(ctx context.Context, login string) (*domain.User, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		}
	}()

	statement, err := tr.PrepareContext(ctx, `
    	SELECT id, login, name, surname, COALESCE(patronymic, '') AS patronymic, role, COALESCE(section, '') AS section,
           student_group, visits, paid, last_scanned, qr_token 
    FROM users 
    WHERE login=$1
	`)

	if err != nil {
		return nil, err
	}
	defer statement.Close()

	var result domain.User
	err = statement.QueryRow(login).Scan(
		&result.ID,
		&result.Login,
		&result.Name,
		&result.Surname,
		&result.Patronymic,
		&result.Role,
		&result.Section,
		&result.StudentGroup,
		&result.Visits,
		&result.Paid,
		&result.Last_scanned,
		&result.QrCode,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrUserNotFound
		}
		return nil, err
	}

	if err := tr.Commit(); err != nil {
		return nil, err
	}

	return &result, nil
}
