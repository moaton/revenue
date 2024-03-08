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
