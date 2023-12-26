package schedule_payment_respositories

import (
	"context"
	schedule_payment_model "personal-budget/schedule/model"
	"personal-budget/util"

	"github.com/google/uuid"
)

// FetchPlan implements SchedulePaymentRepositories.
func (p *schedulePaymentRepositories) FetchPlan(id uuid.UUID) (schedule_payment_model.SchedulePaymentPlan, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `
	Select scheduled_payments.id, scheduled_payments.user_id, scheduled_payments.account_id, 
	scheduled_payments.amount, scheduled_payments.periods,scheduled_payments.paid_period, 
	scheduled_payments.paydate, scheduled_payments.duration, scheduled_payments.is_completed, 
	scheduled_payments.created_at, scheduled_payments.updated_at, json_agg(st.*) from scheduled_payments
	Left Join schedule_transactions st on st.scheduled_payment_id = scheduled_payments.id
	where scheduled_payments.id =$1 Group by scheduled_payments.id
	`
	var model schedule_payment_model.SchedulePaymentPlan
	err := p.Db.QueryRowContext(ctx, stmt, id).
		Scan(&model.ID, &model.UserId, &model.AccountId, &model.Amount, &model.Periods, &model.PaidPeriods, &model.PayDate, &model.Duration, &model.IsCompleted, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}
