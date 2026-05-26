package dto

import "time"

type OrderStatus string

const (
	OrderStatusBooked   OrderStatus = "booked"
	OrderStatusArchived OrderStatus = "archived"
)

type GetOrder struct {
	ID            int64       `json:"id,omitempty"`
	ProfileID     int64       `json:"profile_id,omitempty"`
	BranchID      int64       `json:"branch_id,omitempty"`
	Date          time.Time   `json:"date,omitempty"`
	CustomerPhone string      `json:"customer_phone,omitempty"`
	Description   string      `json:"description,omitempty"`
	Status        OrderStatus `json:"status,omitempty"`
}

type CreateOrder struct {
	ProfileID     int64     `json:"profile_id,omitempty"`
	BranchID      int64     `json:"branch_id,omitempty"`
	Date          time.Time `json:"date,omitempty"`
	CustomerPhone string    `json:"customer_phone,omitempty"`
	Description   string    `json:"description,omitempty"`
}

type SetOrderStatus struct {
	ID     int64       `json:"id,omitempty"`
	Status OrderStatus `json:"status,omitempty"`
}
