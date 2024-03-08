package repository

import (
	"backend/internal/interfaces"
	"backend/internal/repository/revenue"
	"backend/internal/repository/users"
	"backend/pkg/gorm/postgres"
)

type Repository struct {
	db      *postgres.Gorm
	revenue *revenue.Repository
	users   *users.Repository
}

func NewRepository(db *postgres.Gorm) *Repository {
	return &Repository{
		db:      db,
		revenue: revenue.New(),
	}
}

func (r *Repository) Conn() *postgres.Gorm {
	return r.db
}

func (r *Repository) Revenue() interfaces.RevenueRepositoryInterface {
	return r.revenue
}

func (r *Repository) Users() interfaces.UsersRepositoryInterface {
	return r.users
}
