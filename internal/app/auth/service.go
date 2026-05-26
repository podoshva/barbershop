// Package auth
package auth

import (
	"context"
	"fmt"
	"main/internal/adapter/postgres/repos"
	"main/internal/dto"
	"main/pkg"
)

type AuthService struct {
	profileRepository *repos.ProfileRepository
}

func NewAuthService(repo *repos.ProfileRepository) *AuthService {
	return &AuthService{
		profileRepository: repo,
	}
}

func (s *AuthService) Login(ctx context.Context, in dto.LoginIn) (*string, error) {
	out, err := s.profileRepository.GetAuthByLogin(ctx, in.Login)
	if err != nil {
		return nil, err
	}
	if out.Password != in.Password {
		return nil, fmt.Errorf("password incorrect")
	}
	token, err := pkg.GenerateToken(out.ID, out.Role)
	return &token, err
}
