package dto

import (
	"backend/internal/types"

	"github.com/google/uuid"
)

type CreateRevenueRequest struct {
	UserId      string     `json:"user_id"`
	Name        string     `json:"name"`
	Type        types.Type `json:"type"`
	Amount      int64      `json:"amount"`
	Description string     `json:"description"`
	GroupId     uuid.UUID  `json:"group_id"`
}

type CreateRevenueResponse struct {
	Id string `json:"id"`
}
