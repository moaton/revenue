package users

import (
	"backend/internal/entity"
	"backend/pkg/gorm/postgres"
)

type Repository struct {
}

func New() *Repository {
	return &Repository{}
}

func (r *Repository) SignUp(db *postgres.Gorm, req entity.User) error {
	return db.DB.Debug().Create(&req).Error
}

func (r *Repository) Login(db *postgres.Gorm, username string) (entity.User, error) {
	var user entity.User
	err := db.DB.Debug().First(&user, "username = ?", username).Error
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
