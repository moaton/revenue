package usecase

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"time"

	"github.com/google/uuid"
)

func revenueToEntity(req dto.Revenue) (entity.Revenue, error) {
	datetime, err := time.Parse("2006-01-02T15:04:05", req.Datetime)
	if err != nil {
		return entity.Revenue{}, err
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
	}, nil
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
		datetime := revenue.Datetime.Format("2006-01-02 15:04:05")
		revenues = append(revenues, dto.Revenue{
			UserId:      revenue.UserId,
			Name:        revenue.Name,
			Type:        revenue.Type,
			Amount:      revenue.Amount,
			Description: revenue.Description,
			GroupId:     revenue.GroupId,
			Datetime:    datetime,
		})
	}
	return revenues
}
