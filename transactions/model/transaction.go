package model_transaction

import (
	"personal-budget/shared"

	"github.com/google/uuid"
)

type Transaction struct {
	shared.Model
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	UserID    uuid.UUID `json:"user_id"`
	Reference string    `json:"reference"`
	Amount    int       `json:"amount"`
}
