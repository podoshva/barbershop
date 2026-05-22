// Package order
package order

import (
	"context"
	"main/internal/adapter/postgres/repos"
	"main/internal/dto"
)

type OrderService struct {
	orderRepository *repos.OrderRepository
}

func NewOrderService(repo *repos.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: repo,
	}
}

func (s *OrderService) Create(ctx context.Context, in dto.CreateOrder) error {
	return s.orderRepository.Create(ctx, in)
}

func (s *OrderService) Delete(ctx context.Context, id int64) error {
	return s.orderRepository.Delete(ctx, id)
}

func (s *OrderService) Get(ctx context.Context, id int64) (*dto.GetOrder, error) {
	return s.orderRepository.Get(ctx, id)
}
