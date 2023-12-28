package repositories_scheduled_transactions

import (
	"context"
	model_scheduled_transactions "personal-budget/schedule_transactions/model"
	"personal-budget/util"
)

// UpdateUserTransaction implements TransactionRepo.
func (p *scheduledtransactionRepo) UpdateUserTransaction(reference string, status string) (model_scheduled_transactions.ScheduledTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update schedule_transactions  set status = $1 where reference=$2`
	var model model_scheduled_transactions.ScheduledTransaction
	err := p.DB.QueryRowContext(ctx, stmt, status, reference).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// UpdateUserTransaction implements TransactionRepo.
func (p *scheduledtransactionRepo) UpdateUserTransactionType(reference string, typeD string) (model_scheduled_transactions.ScheduledTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update schedule_transactions  set type = $1 where reference=$2`
	var model model_scheduled_transactions.ScheduledTransaction
	err := p.DB.QueryRowContext(ctx, stmt, typeD, reference).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}
