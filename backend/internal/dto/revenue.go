package dto

import (
	"backend/internal/types"

	"github.com/google/uuid"
)

type Revenue struct {
	UserId      string     `json:"user_id"`
	Name        string     `json:"name"`
	Type        types.Type `json:"type"`
	Amount      int64      `json:"amount"`
	Description *string    `json:"description"`
	GroupId     *uuid.UUID `json:"group_id"`
	Datetime    string     `json:"datetime"`
}

type CreateRevenueResponse struct {
	Id string `json:"id"`
}

type GetRevenuesRequest struct {
	Pagination Pagination `json:"pagination"`
}
type GetRevenuesResponse struct {
	Revenues []Revenue `json:"revenues"`
	Total    int64     `json:"total"`
}
