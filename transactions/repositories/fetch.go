package repositories_transaction

import (
	"context"
	model_transaction "personal-budget/transactions/model"
	"personal-budget/util"

	"github.com/google/uuid"
)

// GetUserTransaction implements TransactionRepo.
func (p *transactionRepo) GetUserTransaction(userId uuid.UUID) (model_transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, type, status, user_id, reference, amount, created_at, updated_at from transactions where user_id=$1 limit 1 DESC created_at`
	var model model_transaction.Transaction
	err := p.DB.QueryRowContext(ctx, stmt, userId).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// GetUserTransactionByReference implements TransactionRepo.
func (p *transactionRepo) GetUserTransactionByReference(reference string) (model_transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, type, status, user_id, reference, amount, created_at, updated_at from transactions where reference=$1`
	var model model_transaction.Transaction
	err := p.DB.QueryRowContext(ctx, stmt, reference).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// GetUserTransactions implements TransactionRepo.
func (p *transactionRepo) GetUserTransactions(userId uuid.UUID) ([]model_transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, type, status, user_id, reference, amount, created_at, updated_at from transactions where reference=$1`
	rows, err := p.DB.QueryContext(ctx, stmt, userId)
	results := make([]model_transaction.Transaction, 0)
	for rows.Next() {
		var model model_transaction.Transaction
		rows.Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
		results = append(results, model)
	}
	return results, err
}
