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

func (r *revenue) CreateRevenue(ctx context.Context, req dto.Revenue) (string, error) {
	revenue := revenueToEntity(req)
	err := r.db.Revenue().CreateRevenue(r.db.Conn(), revenue)
	if err != nil {
		return "", err
	}
	return revenue.ID.String(), nil
}

func (r *revenue) GetRevenues(ctx context.Context, req dto.GetRevenuesRequest) ([]dto.Revenue, int64, error) {
	req.Pagination.Validate()
	revenues, total, err := r.db.Revenue().GetRevenues(r.db.Conn(), getRevenuesRequestToEntity(req))
	if err != nil {
		return nil, 0, err
	}
	return convertRevenues(revenues), total, nil
}
