package usecase

import (
	"backend/internal/dto"
	"backend/internal/interfaces"
	"context"
)

type revenue struct {
	db interfaces.Repository
}

func newRevenue(repo interfaces.Repository) *revenue {
	return &revenue{
		db: repo,
	}
}

func (r *revenue) CreateRevenue(ctx context.Context, req dto.CreateRevenueRequest) (string, error) {
	revenue := revenueToEntity(req)
	err := r.db.Revenue().CreateRevenue(r.db.Conn(), revenue)
	if err != nil {
		return "", err
	}
	return revenue.ID.String(), nil
}
