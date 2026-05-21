// Package branch
package branch

import (
	"context"
	"main/internal/adapter/postgres/repos"
)

type BranchService struct {
	branchRepository *repos.BranchRepository
}

func NewBranchService(branchRepo *repos.BranchRepository) *BranchService {
	return &BranchService{
		branchRepository: branchRepo,
	}
}

func (s *BranchService) Create(ctx context.Context, name string) error {
	err := s.branchRepository.Create(ctx, repos.CreateBranch{Name: name})
	return err
}

func (s *BranchService) Delete(ctx context.Context, id int64) error {
	err := s.branchRepository.Delete(ctx, id)
	return err
}
