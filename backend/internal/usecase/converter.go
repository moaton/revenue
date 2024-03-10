package usecase

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"time"

	"github.com/google/uuid"
)

func revenueToEntity(req dto.Revenue) entity.Revenue {
	now := time.Now().UTC()
	datetime := now
	if req.Datetime != nil {
		datetime = *req.Datetime
	}
	return entity.Revenue{
		ID:          uuid.New(),
		UserId:      req.UserId,
		Name:        req.Name,
		Type:        req.Type,
		Amount:      req.Amount,
		Description: req.Description,
		GroupId:     req.GroupId,
		Datetime:    datetime,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}

func getRevenuesRequestToEntity(req dto.GetRevenuesRequest) entity.RevenueFilter {
	return entity.RevenueFilter{
		Pagination: entity.Pagination{
			Limit:  int(req.Pagination.Size),
			Offset: int((req.Pagination.Page - 1) * req.Pagination.Size),
		},
	}
}

func convertRevenues(req []entity.Revenue) []dto.Revenue {
	revenues := make([]dto.Revenue, 0, len(req))
	for _, revenue := range req {
		revenues = append(revenues, dto.Revenue{
			UserId:      revenue.UserId,
			Name:        revenue.Name,
			Type:        revenue.Type,
			Amount:      revenue.Amount,
			Description: revenue.Description,
			GroupId:     revenue.GroupId,
			Datetime:    &revenue.Datetime,
		})
	}
	return revenues
}
