package model_account

import (
	"personal-budget/shared"

	"github.com/google/uuid"
)

type Account struct {
	shared.Model
	Name     string    `json:"name" validate:"required"`
	BankCode string    `json:"bank_code" validate:"required"`
	Number   string    `json:"number" validate:"required"`
	UserId   uuid.UUID `json:"user_id"`
}

type AccountParam struct {
	ID string `uri:"id" binding:"required"`
}
