package schedule_payment_model

import (
	"encoding/json"
	"personal-budget/shared"
	"time"

	"github.com/google/uuid"
)

type SchedulePayment struct {
	shared.Model
	UserId      uuid.UUID `json:"user_id"`
	AccountId   uuid.UUID `json:"account_id"`
	Amount      int       `json:"amount"`
	Periods     int       `json:"periods"`
	Duration    int       `json:"duration"` ///// duration in days and its the time frame to make the payment
	PaidPeriods int       `json:"paid_periods"`
	PayDate     time.Time `json:"pay_date"`
	IsCompleted bool      `json:"is_completed"`
}

type SchedulePaymentPlan struct {
	shared.Model
	UserId       uuid.UUID       `json:"user_id"`
	AccountId    uuid.UUID       `json:"account_id"`
	Amount       int             `json:"account"`
	Periods      int             `json:"periods"`
	Duration     int             `json:"duration"` ///// duration in days and its the time frame to make the payment
	PaidPeriods  int             `json:"paid_periods"`
	PayDate      time.Time       `json:"pay_date"`
	IsCompleted  bool            `json:"is_completed"`
	Transactions json.RawMessage `json:"transactions"`
	// Transactions []model_scheduled_transactions.ScheduledTransaction `json:"transactions"`
}

type SchedulePaymentAccount struct {
	SchedulePayment
	RecipientCode string `json:"recipient_code"`
}

// recipient_code
