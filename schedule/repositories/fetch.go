package schedule_payment_respositories

import (
	"context"
	"database/sql"
	schedule_payment_model "personal-budget/schedule/model"
	"personal-budget/shared"
	"personal-budget/util"

	"github.com/google/uuid"
)

// FetchPlan implements SchedulePaymentRepositories.
func (p *schedulePaymentRepositories) FetchPlan(id uuid.UUID) (schedule_payment_model.SchedulePaymentPlan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `
	Select scheduled_payments.id, scheduled_payments.user_id, scheduled_payments.account_id, 
	scheduled_payments.amount, scheduled_payments.periods,scheduled_payments.paid_periods, 
	scheduled_payments.paydate, scheduled_payments.duration, scheduled_payments.is_completed, 
	scheduled_payments.created_at, scheduled_payments.updated_at, json_agg(st.*) as transactions from scheduled_payments
	Left Join schedule_transactions st on st.scheduled_payment_id = scheduled_payments.id
	where scheduled_payments.id =$1 Group by scheduled_payments.id
	`
	var model schedule_payment_model.SchedulePaymentPlan
	err := p.Db.QueryRowContext(ctx, stmt, id).
		Scan(&model.ID, &model.UserId, &model.AccountId, &model.Amount, &model.Periods, &model.PaidPeriods, &model.PayDate, &model.Duration, &model.IsCompleted, &model.CreatedAt, &model.UpdatedAt, &model.Transactions)
	if len(model.Transactions) < 8 {
		model.Transactions = shared.JsonRawMessage
	}
	return model, err
}

func (p *schedulePaymentRepositories) FetchPlansRecords(batchSize int, processedRows int, tx *sql.Tx) ([]schedule_payment_model.SchedulePayment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `Select id, user_id, account_id ,amount,periods,paid_periods,paydate,duration, is_completed,created_at, updated_at from scheduled_payments where is_completed <> true and (paydate  BETWEEN NOW() - INTERVAL '3 hours' AND NOW()+ INTERVAL '3 hours') 
	 limit $1 offset $2 `
	rows, err := tx.QueryContext(ctx, stmt, batchSize, processedRows)
	defer rows.Close()
	var models []schedule_payment_model.SchedulePayment
	for rows.Next() {
		var model schedule_payment_model.SchedulePayment
		rows.Scan(&model.ID, &model.UserId, &model.AccountId, &model.Amount, &model.Periods, &model.PaidPeriods, &model.PayDate, &model.Duration, &model.IsCompleted, &model.CreatedAt, &model.UpdatedAt)
		models = append(models, model)
	}
	return models, err
}

func (p *schedulePaymentRepositories) FetchPlansRecordsCounter() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	var totalRow int
	stmt := `Select count(*) from scheduled_payments where is_completed <> true and (paydate  BETWEEN NOW() - INTERVAL '3 hours' AND NOW()+ INTERVAL '3 hours')`
	err := p.Db.QueryRowContext(ctx, stmt).
		Scan(&totalRow)

	return totalRow, err
}
