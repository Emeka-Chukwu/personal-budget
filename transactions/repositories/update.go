package repositories_transaction

import (
	"context"
	model_transaction "personal-budget/transactions/model"
	"personal-budget/util"
)

// UpdateUserTransaction implements TransactionRepo.
func (p *transactionRepo) UpdateUserTransaction(reference string, status string) (model_transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update transactions set status = $1 where reference=$2`
	var model model_transaction.Transaction
	err := p.DB.QueryRowContext(ctx, stmt, status, reference).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// UpdateUserTransaction implements TransactionRepo.
func (p *transactionRepo) UpdateUserTransactionType(reference string, typeD string) (model_transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update transactions set type = $1 where reference=$2`
	var model model_transaction.Transaction
	err := p.DB.QueryRowContext(ctx, stmt, typeD, reference).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}
