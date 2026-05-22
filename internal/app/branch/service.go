// Package branch
package branch

import (
	"context"
	"main/internal/adapter/postgres/repos"
	"main/internal/dto"
)

type BranchService struct {
	branchRepository *repos.BranchRepository
}

func NewBranchService(branchRepo *repos.BranchRepository) *BranchService {
	return &BranchService{
		branchRepository: branchRepo,
	}
}

func (s *BranchService) Create(ctx context.Context, dto dto.CreateBranch) error {
	return s.branchRepository.Create(ctx, dto)
}

func (s *BranchService) Delete(ctx context.Context, id int64) error {
	return s.branchRepository.Delete(ctx, id)
}

func (s *BranchService) Get(ctx context.Context, id int64) (*dto.GetBranch, error) {
	return s.branchRepository.Get(ctx, id)
}

func (s *BranchService) GetAll(ctx context.Context) ([]dto.GetBranch, error) {
	return s.branchRepository.GetAll(ctx)
}
