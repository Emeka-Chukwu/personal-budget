package schedule_payment_respositories

import (
	"context"
	"database/sql"
	schedule_payment_model "personal-budget/schedule/model"
	"personal-budget/util"
)

// CreatePlan implements SchedulePaymentRepositories.
func (p *schedulePaymentRepositories) CreatePlan(req schedule_payment_model.SchedulePayment) (schedule_payment_model.SchedulePayment, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into scheduled_payments (user_id, account_id, amount, periods,paid_periods, paydate, duration) values ($1, $2, $3, $4, $5, $6, $7)
	returning id, user_id, account_id, amount, periods,paid_periods, paydate, duration, is_completed, created_at, updated_at`
	var model schedule_payment_model.SchedulePayment
	err := p.Db.QueryRowContext(ctx, stmt, req.UserId, req.AccountId, req.Amount, req.Periods, req.PaidPeriods, req.PayDate, req.Duration).
		Scan(&model.ID, &model.UserId, &model.AccountId, &model.Amount, &model.Periods, &model.PaidPeriods, &model.PayDate, &model.Duration, &model.IsCompleted, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// CreatePlanTx implements SchedulePaymentRepositories.
func (p *schedulePaymentRepositories) CreatePlanTx(req schedule_payment_model.SchedulePayment, tx *sql.Tx) (schedule_payment_model.SchedulePayment, *sql.Tx, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into scheduled_payments (user_id, account_id, amount, periods,paid_periods, paydate, duration) values ($1, $2, $3, $4, $5, $6, $7)
	returning id, user_id, account_id, amount, periods,paid_periods, paydate, duration, is_completed, created_at, updated_at`
	var model schedule_payment_model.SchedulePayment
	err := tx.QueryRowContext(ctx, stmt, req.UserId, req.AccountId, req.Amount, req.Periods, req.PaidPeriods, req.PayDate, req.Duration).
		Scan(&model.ID, &model.UserId, &model.AccountId, &model.Amount, &model.Periods, &model.PaidPeriods, &model.PayDate, &model.Duration, &model.IsCompleted, &model.CreatedAt, &model.UpdatedAt)
	return model, tx, err
}
