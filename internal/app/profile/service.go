// Package profile
package profile

import (
	"context"
	"main/internal/adapter/postgres/repos"
	"main/internal/dto"
)

type ProfileService struct {
	profileRepository *repos.ProfileRepository
}

func NewProfileService(repo *repos.ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepository: repo,
	}
}

func (s *ProfileService) Create(ctx context.Context, in dto.CreateProfile) error {
	return s.profileRepository.Create(ctx, in)
}

func (s *ProfileService) Delete(ctx context.Context, id int64) error {
	return s.profileRepository.Delete(ctx, id)
}

func (s *ProfileService) Get(ctx context.Context, id int64) (*dto.GetProfile, error) {
	return s.profileRepository.Get(ctx, id)
}

func (s *ProfileService) GetAll(ctx context.Context) ([]dto.GetProfile, error) {
	return s.profileRepository.GetAll(ctx)
}
