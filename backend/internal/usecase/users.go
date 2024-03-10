package usecase

import (
	"backend/internal/dto"
	"backend/internal/entity"
	"backend/internal/interfaces"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type users struct {
	db interfaces.Repository
}

func newUsers(repo interfaces.Repository) *users {
	return &users{
		db: repo,
	}
}

func (u *users) SignUp(ctx context.Context, req dto.SignUpRequest) error {
	user := entity.User{
		Id:        uuid.New(),
		Username:  req.Username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return fmt.Errorf("password hash error %v", err)
	}
	user.Password = string(hash)
	err = u.db.Users().SignUp(u.db.Conn(), user)
	if err != nil {
		return err
	}
	return nil
}

func (u *users) Login(ctx context.Context, req dto.LoginRequest) (string, error) {
	user, err := u.db.Users().Login(u.db.Conn(), req.Username)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("password hash error %v", err)
	}

	ttl := time.Hour * 5

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"exp":    time.Now().Add(ttl).Unix(),
	})
	tokenString, err := token.SignedString([]byte("#3ad$tasd"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
