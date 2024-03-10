package revenue

import (
	"backend/internal/entity"
	"backend/pkg/gorm/postgres"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) CreateRevenue(db *postgres.Gorm, req entity.Revenue) error {
	return db.DB.Debug().Create(&req).Error
}

func (r *Repository) GetRevenues(db *postgres.Gorm, req entity.RevenueFilter) ([]entity.Revenue, int64, error) {
	var revenues []entity.Revenue
	var total int64
	err := db.DB.Debug().Table("revenues").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.DB.Debug().Table("revenues").Limit(req.Pagination.Limit).Offset(req.Pagination.Offset).Find(&revenues).Error
	if err != nil {
		return nil, 0, err
	}
	return revenues, total, nil
}
