package schedule_payment_model

import (
	"personal-budget/shared"
	"time"

	model_scheduled_transactions "personal-budget/schedule_transactions/model"

	"github.com/google/uuid"
)

type SchedulePayment struct {
	shared.Model
	UserId      uuid.UUID `json:"user_id"`
	AccountId   uuid.UUID `json:"account_id"`
	Amount      int       `json:"account"`
	Periods     int       `json:"periods"`
	Duration    int       `json:"duration"` ///// duration in days and its the time frame to make the payment
	PaidPeriods int       `json:"paid_periods"`
	PayDate     time.Time `json:"pay_date"`
	IsCompleted bool      `json:"is_completed"`
}

type SchedulePaymentPlan struct {
	shared.Model
	UserId       uuid.UUID                                           `json:"user_id"`
	AccountId    uuid.UUID                                           `json:"account_id"`
	Amount       int                                                 `json:"account"`
	Periods      int                                                 `json:"periods"`
	Duration     int                                                 `json:"duration"` ///// duration in days and its the time frame to make the payment
	PaidPeriods  int                                                 `json:"paid_periods"`
	PayDate      time.Time                                           `json:"pay_date"`
	IsCompleted  bool                                                `json:"is_completed"`
	Transactions []model_scheduled_transactions.ScheduledTransaction `json:"transactions"`
}
