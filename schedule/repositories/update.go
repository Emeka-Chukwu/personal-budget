package schedule_payment_respositories

import (
	"context"
	"database/sql"
	schedule_payment_model "personal-budget/schedule/model"
	"personal-budget/util"
	"time"
)

// UpdatePlanTx implements SchedulePaymentRepositories.
func (*schedulePaymentRepositories) UpdatePlanTx(req schedule_payment_model.SchedulePayment, tx *sql.Tx) (schedule_payment_model.SchedulePayment, *sql.Tx, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `UPDATE scheduled_payments set paid_periods= $1, paydate= $2, is_completed=$3, updated_at= $4 where is_completed=$5 and id=$6 
	returning id, user_id, account_id ,amount,periods,paid_periods,paydate,duration, is_completed,created_at, updated_at`
	var model schedule_payment_model.SchedulePayment
	err := tx.QueryRowContext(ctx, stmt, req.PaidPeriods, req.PayDate, req.IsCompleted, time.Now(), req.IsCompleted, req.ID).
		Scan(&model.ID, &model.UserId, &model.AccountId, &model.Amount, &model.Periods, &model.PaidPeriods, &model.PayDate, &model.Duration, &model.IsCompleted, &model.CreatedAt, &model.UpdatedAt)
	return model, tx, err
}
