package repos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BranchRepo struct {
	pool *pgxpool.Pool
}

func NewBranchRepo(pool *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{
		pool: pool,
	}
}

type CreateBranch struct {
	Name string
}

func (r *BranchRepo) CreateBranch(ctx context.Context, in CreateBranch) error {
	if _, err := r.pool.Exec(ctx, "INSERT INTO branches (name) VALUES ($1)", in.Name); err != nil {
		return fmt.Errorf("create branch: %w", err)
	}
	return nil
}
