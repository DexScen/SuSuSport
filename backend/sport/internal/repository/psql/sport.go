package psql

import (
	"context"
	"database/sql"

	//"github.com/DexScen/SuSuSport/backend/auth/internal/domain"
	//"github.com/DexScen/SuSuSport/backend/sport/internal/errors"
	"github.com/DexScen/SuSuSport/backend/sport/internal/domain"
	_ "github.com/lib/pq"
)

type Sport struct {
	db *sql.DB
}

func NewSport(db *sql.DB) *Sport {
	return &Sport{db: db}
}

func (s *Sport) GetSections(ctx context.Context) (*[]string, error) {
	tr, err := s.db.Begin()
	if err != nil {
		return nil, err
	}

	statement, err := tr.PrepareContext(ctx, `
		SELECT name
		FROM sections
		ORDER BY name ASC
	`)
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	defer statement.Close()

	rows, err := statement.QueryContext(ctx)
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	defer rows.Close()

	var sections []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			tr.Rollback()
			return nil, err
		}
		sections = append(sections, name)
	}

	if err := rows.Err(); err != nil {
		tr.Rollback()
		return nil, err
	}

	if err := tr.Commit(); err != nil {
		return nil, err
	}

	return &sections, nil
}

func (s *Sport) GetSectionInfoByName(ctx context.Context, name string) (*domain.Section, error){
	query := "SELECT * FROM sections WHERE name = $1"
	row := s.db.QueryRow(query, name)
	section := domain.Section{}
	err := row.Scan(&section.ID, &section.Name, &section.Info, &section.Schedule)
	if err != nil {
		return nil, err
	}
	return  &section, nil
}