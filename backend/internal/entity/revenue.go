package entity

import (
	"go/types"

	"github.com/google/uuid"
)

type Revenue struct {
	ID          uuid.UUID  `json:"id"`
	UserId      string     `json:"user_id"`
	Name        string     `json:"name"`
	Type        types.Type `json:"type"`
	Amount      int64      `json:"amount"`
	Description string     `json:"description"`
	GroupId     uuid.UUID  `json:"group_id"`
}
