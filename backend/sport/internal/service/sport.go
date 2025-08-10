package service

import (
	"context"

	"github.com/DexScen/SuSuSport/backend/sport/internal/domain"
	//"errors"
	//e "github.com/DexScen/SuSuSport/backend/sport/internal/errors"
)

type SportRepository interface {
	GetSections (ctx context.Context) (*[]string, error)
	GetSectionInfoByName(ctx context.Context, name string) (*domain.Section, error)
}

type Sport struct {
	repo SportRepository
}

func NewSport(repo SportRepository) *Sport {
	return &Sport{
		repo: repo,
	}
}

func (s *Sport) GetSections(ctx context.Context) (*[]string, error) {
	return s.repo.GetSections(ctx)
}

func (s *Sport) GetSectionInfoByName(ctx context.Context, name string) (*domain.Section, error) {
	return s.repo.GetSectionInfoByName(ctx, name)
}