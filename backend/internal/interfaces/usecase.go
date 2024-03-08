package interfaces

import (
	"backend/internal/dto"
	"context"
)

type Usecase interface {
	Users() UsersUsecaseInterface
	Revenue() RevenueUsecaseInterface
}

type UsersUsecaseInterface interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) error
	Login(ctx context.Context, req dto.LoginRequest) (string, error)
}

type RevenueUsecaseInterface interface {
	CreateRevenue(ctx context.Context, req dto.CreateRevenueRequest) (string, error)
}
