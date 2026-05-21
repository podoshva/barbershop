// Package profile
package profile

import (
	"context"
	"main/internal/adapter/postgres/repos"
)

type ProfileService struct {
	profileRepository *repos.ProfileRepository
}

func NewProfileService(repo *repos.ProfileRepository) *ProfileService {
	return &ProfileService{
		profileRepository: repo,
	}
}

func (s *ProfileService) Create(ctx context.Context, in repos.CreateProfile) error {
	return s.profileRepository.Create(ctx, in)
}

func (s *ProfileService) Delete(ctx context.Context, id int64) error {
	return s.profileRepository.Delete(ctx, id)
}
