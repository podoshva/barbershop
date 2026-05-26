package handlers

import (
	"main/internal/app/auth"
	"main/internal/app/branch"
	"main/internal/app/order"
	"main/internal/app/profile"
)

type Handlers struct {
	BranchService  *branch.BranchService
	ProfileService *profile.ProfileService
	OrderService   *order.OrderService
	AuthService    *auth.AuthService
}
