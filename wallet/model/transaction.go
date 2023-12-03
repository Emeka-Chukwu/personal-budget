package wallet_model

import (
	"personal-budget/shared"

	"github.com/google/uuid"
)

type Transaction struct {
	shared.Model
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	Amount    int       `json:"amount"`
	UserID    uuid.UUID `json:"user_id"`
	Reference string    `json:"reference"`
}
