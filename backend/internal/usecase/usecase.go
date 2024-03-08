package usecase

import (
	"backend/internal/interfaces"
)

type Usecase struct {
	users   *users
	revenue *revenue
	repo    interfaces.Repository
}

func New(repo interfaces.Repository) *Usecase {
	return &Usecase{
		users:   newUsers(repo),
		revenue: newRevenue(repo),
		repo:    repo,
	}
}

func (u *Usecase) Users() interfaces.UsersUsecaseInterface {
	return u.users
}

func (u *Usecase) Revenue() interfaces.RevenueUsecaseInterface {
	return u.revenue
}
