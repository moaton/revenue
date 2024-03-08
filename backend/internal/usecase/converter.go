package usecase

import (
	"backend/internal/dto"
	"backend/internal/entity"

	"github.com/google/uuid"
)

func revenueToEntity(req dto.CreateRevenueRequest) entity.Revenue {
	return entity.Revenue{
		ID: uuid.New(),
	}
}
