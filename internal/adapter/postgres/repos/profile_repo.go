package repos

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	pool *pgxpool.Pool
}

func NewProfileRepository(pool *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{
		pool: pool,
	}
}

type CreateProfile struct {
	BranchID int64
	FullName string
	Login    string
	Password string
	Role     string
}

func (r *ProfileRepository) Create(ctx context.Context, in CreateProfile) error {
	if _, err := r.pool.Exec(ctx, "INSERT INTO profiles (branch_id, full_name, login, password, role) VALUES ($1, $2, $3, $4, $5)", in.BranchID, in.FullName, in.Login, in.Password, in.Role); err != nil {
		return fmt.Errorf("create profile: %w", err)
	}
	return nil
}

func (r *ProfileRepository) Delete(ctx context.Context, id int64) error {
	if _, err := r.pool.Exec(ctx, "DELETE FROM profiles WHERE id = $1", id); err != nil {
		return fmt.Errorf("delete profile: %w", err)
	}
	return nil
}
