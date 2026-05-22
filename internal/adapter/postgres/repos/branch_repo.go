// Package repos
package repos

import (
	"context"
	"fmt"
	"main/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BranchRepository struct {
	pool *pgxpool.Pool
}

func NewBranchRepository(pool *pgxpool.Pool) *BranchRepository {
	return &BranchRepository{
		pool: pool,
	}
}

func (r *BranchRepository) Create(ctx context.Context, in dto.CreateBranch) error {
	if _, err := r.pool.Exec(ctx, "INSERT INTO branches (name) VALUES ($1)", in.Name); err != nil {
		return fmt.Errorf("create branch: %w", err)
	}
	return nil
}

func (r *BranchRepository) Delete(ctx context.Context, id int64) error {
	if _, err := r.pool.Exec(ctx, "DELETE FROM branches WHERE id = $1", id); err != nil {
		return fmt.Errorf("delete branch: %w", err)
	}
	return nil
}

func (r *BranchRepository) Get(ctx context.Context, id int64) (*dto.GetBranch, error) {
	var branch dto.GetBranch
	err := r.pool.QueryRow(
		ctx,
		"SELECT id, name FROM branches WHERE id = $1",
		id,
	).Scan(
		&branch.ID,
		&branch.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("get branch: %w", err)
	}
	return &branch, nil
}
