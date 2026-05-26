package repos

import (
	"context"
	"fmt"
	"main/internal/dto"

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

func (r *ProfileRepository) Create(ctx context.Context, in dto.CreateProfile) error {
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

func (r *ProfileRepository) Get(ctx context.Context, id int64) (*dto.GetProfile, error) {
	var profile dto.GetProfile
	err := r.pool.QueryRow(ctx, "SELECT id, branch_id, full_name, login, role FROM profiles WHERE id = $1", id).Scan(
		&profile.ID,
		&profile.BranchID,
		&profile.FullName,
		&profile.Login,
		&profile.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("get profile: %w", err)
	}
	return &profile, nil
}

func (r *ProfileRepository) GetAll(ctx context.Context) ([]dto.GetProfile, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, branch_id, full_name, login, role FROM profiles")
	if err != nil {
		return nil, fmt.Errorf("get profiles: %w", err)
	}
	defer rows.Close()
	var profiles []dto.GetProfile
	for rows.Next() {
		var profile dto.GetProfile
		err := rows.Scan(
			&profile.ID,
			&profile.BranchID,
			&profile.FullName,
			&profile.Login,
			&profile.Role,
		)
		if err != nil {
			return nil, fmt.Errorf("scan profile: %w", err)
		}
		profiles = append(profiles, profile)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return profiles, nil
}

func (r *ProfileRepository) GetAuthByLogin(ctx context.Context, login string) (*dto.LoginOut, error) {
	var out dto.LoginOut
	err := r.pool.QueryRow(ctx, "SELECT id, password, role FROM profiles WHERE login = $1", login).Scan(&out.ID, &out.Password, &out.Role)
	if err != nil {
		return nil, fmt.Errorf("get auth by login error: %w", err)
	}
	return &out, nil
}
