package wallet_model

import (
	"personal-budget/shared"

	"github.com/google/uuid"
)

type Wallet struct {
	shared.Model
	UserID  uuid.UUID `json:"user_id"`
	Balance int       `json:"balance" validate:"required"`
}
