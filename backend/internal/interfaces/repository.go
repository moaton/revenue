package interfaces

import (
	"backend/internal/entity"
	"backend/pkg/gorm/postgres"
)

type Repository interface {
	Conn() *postgres.Gorm
	Users() UsersRepositoryInterface
	Revenue() RevenueRepositoryInterface
}

type UsersRepositoryInterface interface {
	SignUp(db *postgres.Gorm, req entity.User) error
	Login(db *postgres.Gorm, username string) (entity.User, error)
}

type RevenueRepositoryInterface interface {
	CreateRevenue(db *postgres.Gorm, req entity.Revenue) error
	GetRevenues(db *postgres.Gorm, req entity.RevenueFilter) ([]entity.Revenue, int64, error)
	GetCharts(db *postgres.Gorm) ([]int64, error)
}
