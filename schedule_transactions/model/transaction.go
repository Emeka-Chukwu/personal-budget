package model_scheduled_transactions

import (
	"personal-budget/shared"

	"github.com/google/uuid"
)

type ScheduledTransaction struct {
	shared.Model
	Type              string    `json:"type"`
	Status            string    `json:"status"`
	UserID            uuid.UUID `json:"user_id"`
	SchedulePaymentID uuid.UUID `json:"scheduled_payment_id"`
	PaidPeriod        int       `json:"paid_period"`
	Reference         string    `json:"reference"`
	Amount            int       `json:"amount"`
}
