package schedule_payment_respositories

import (
	"database/sql"
	schedule_payment_model "personal-budget/schedule/model"

	"github.com/google/uuid"
)

type SchedulePaymentRepositories interface {
	CreatePlan(model schedule_payment_model.SchedulePayment) (schedule_payment_model.SchedulePayment, error)
	FetchPlan(id uuid.UUID) (schedule_payment_model.SchedulePaymentPlan, error)
	ListPlan(userId uuid.UUID) ([]schedule_payment_model.SchedulePaymentPlan, error)
	CreatePlanTx(req schedule_payment_model.SchedulePayment, tx *sql.Tx) (schedule_payment_model.SchedulePayment, *sql.Tx, error)
	UpdatePlanTx(model schedule_payment_model.SchedulePayment, tx *sql.Tx) (schedule_payment_model.SchedulePayment, *sql.Tx, error)
	FetchPlansRecords(batchSize int, processedRows int, tx *sql.Tx) ([]schedule_payment_model.SchedulePayment, error)
	FetchPlansRecordsCounter() (int, error)
}
type schedulePaymentRepositories struct {
	Db *sql.DB
}

func NewSchedulePaymentRepositories(Db *sql.DB) SchedulePaymentRepositories {
	return &schedulePaymentRepositories{Db: Db}
}
