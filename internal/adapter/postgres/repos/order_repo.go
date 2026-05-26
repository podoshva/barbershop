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
	if _, err := r.pool.Exec(ctx, "INSERT INTO orders (profile_id, branch_id, date, customer_phone, description) VALUES ($1, $2, $3, $4, $5)", in.ProfileID, in.BranchID, in.Date, in.CustomerPhone, in.Description); err != nil {
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

func (r *OrderRepository) GetAll(ctx context.Context) ([]dto.GetOrder, error) {
	rows, err := r.pool.Query(ctx, "SELECT id, profile_id, branch_id, date, customer_phone, description, status FROM orders")
	if err != nil {
		return nil, fmt.Errorf("get orders: %w", err)
	}
	defer rows.Close()
	var orders []dto.GetOrder
	for rows.Next() {
		var order dto.GetOrder
		err := rows.Scan(
			&order.ID,
			&order.ProfileID,
			&order.BranchID,
			&order.Date,
			&order.CustomerPhone,
			&order.Description,
			&order.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("scan order: %w", err)
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}
	return orders, nil
}

func (r *OrderRepository) SetStatus(ctx context.Context, in dto.SetOrderStatus) (*dto.GetOrder, error) {
	var order dto.GetOrder
	err := r.pool.QueryRow(ctx, "UPDATE orders SET status = $1 WHERE id = $2 RETURNING id, profile_id, branch_id, date, customer_phone, description, status", in.Status, in.ID).Scan(
		&order.ID,
		&order.ProfileID,
		&order.BranchID,
		&order.Date,
		&order.CustomerPhone,
		&order.Description,
		&order.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("set order status: %w", err)
	}
	return &order, nil
}
