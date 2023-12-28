package repositories_scheduled_transactions

import (
	"context"
	model_scheduled_transactions "personal-budget/schedule_transactions/model"
	"personal-budget/util"

	"github.com/google/uuid"
)

// GetUserTransaction implements TransactionRepo.
func (p *scheduledtransactionRepo) GetUserTransaction(userId uuid.UUID) (model_scheduled_transactions.ScheduledTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, type, status, user_id, reference, amount, created_at, updated_at from transactions where user_id=$1 limit 1 DESC created_at`
	var model model_scheduled_transactions.ScheduledTransaction
	err := p.DB.QueryRowContext(ctx, stmt, userId).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// GetUserTransactionByReference implements TransactionRepo.
func (p *scheduledtransactionRepo) GetUserTransactionByReference(reference string) (model_scheduled_transactions.ScheduledTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, type, status, user_id, reference, amount,paid_period, created_at, updated_at from transactions where reference=$1`
	var model model_scheduled_transactions.ScheduledTransaction
	err := p.DB.QueryRowContext(ctx, stmt, reference).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.PaidPeriod, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// GetUserTransactions implements TransactionRepo.
func (p *scheduledtransactionRepo) GetUserTransactions(userId uuid.UUID) ([]model_scheduled_transactions.ScheduledTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, type, status, user_id, reference, amount, paid_periods, created_at, updated_at from schedule_transactions where reference=$1`
	rows, err := p.DB.QueryContext(ctx, stmt, userId)
	results := make([]model_scheduled_transactions.ScheduledTransaction, 0)
	for rows.Next() {
		var model model_scheduled_transactions.ScheduledTransaction
		rows.Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.PaidPeriod, &model.CreatedAt, &model.UpdatedAt)
		results = append(results, model)
	}
	return results, err
}

func (*scheduledtransactionRepo) GetPlanScheduledTransaction(userId uuid.UUID) (model_scheduled_transactions.ScheduledTransaction, error) {
	panic("unimplemented")
}
