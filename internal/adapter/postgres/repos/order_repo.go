package repos

import (
	"context"
	"fmt"
	"main/internal/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	pool *pgxpool.Pool
}

func NewOrderRepository(pool *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{
		pool: pool,
	}
}

func (r *OrderRepository) Create(ctx context.Context, in dto.CreateOrder) error {
	if _, err := r.pool.Exec(ctx, "INSERT INTO orders (profile_id, branch_id, date, customer_phone, description, status) VALUES ($1, $2, $3, $4, $5, $6)", in.ProfileID, in.BranchID, in.Date, in.CustomerPhone, in.Description, in.Status); err != nil {
		return fmt.Errorf("create order: %w", err)
	}
	return nil
}

func (r *OrderRepository) Delete(ctx context.Context, id int64) error {
	if _, err := r.pool.Exec(ctx, "DELETE FROM orders WHERE id = $1", id); err != nil {
		return fmt.Errorf("delete order: %w", err)
	}
	return nil
}

func (r *OrderRepository) Get(ctx context.Context, id int64) (*dto.GetOrder, error) {
	var order dto.GetOrder
	err := r.pool.QueryRow(ctx, "SELECT id, profile_id, branch_id, date, customer_phone, description, status FROM orders WHERE id = $1", id).Scan(
		&order.ID,
		&order.ProfileID,
		&order.BranchID,
		&order.Date,
		&order.CustomerPhone,
		&order.Description,
		&order.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("get order: %w", err)
	}
	return &order, nil
}
