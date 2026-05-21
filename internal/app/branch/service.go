package branch

import (
	"context"
	"main/internal/adapter/postgres/repos"
)

type BranchService struct {
	branchRepo *repos.BranchRepo
}

func NewBranchService(branchRepo *repos.BranchRepo) *BranchService {
	return &BranchService{
		branchRepo: branchRepo,
	}
}

func (s *BranchService) CreateBranch(ctx context.Context, name string) error {
	err := s.branchRepo.CreateBranch(ctx, repos.CreateBranch{Name: name})
	return err
}
