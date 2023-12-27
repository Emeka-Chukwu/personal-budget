package repositories_scheduled_transactions

import (
	"context"
	"database/sql"
	model_scheduled_transactions "personal-budget/schedule_transactions/model"
	"personal-budget/util"
)

// CreateUserTransaction implements TransactionRepo.
func (p *scheduledtransactionRepo) CreateUserTransaction(req model_scheduled_transactions.ScheduledTransaction) (model_scheduled_transactions.ScheduledTransaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into schedule_transactions (type, status, user_id, reference, amount,paid_period) values ($1, $2, $3, $4, $5, $6)
	returning id, type, status, user_id, reference, amount, paid_period, created_at, updated_at`
	var model model_scheduled_transactions.ScheduledTransaction
	err := p.DB.QueryRowContext(ctx, stmt, req.Type, req.Status, req.UserID, req.Reference, req.Amount, req.PaidPeriod).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.PaidPeriod, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// CreateUserTransaction implements TransactionRepo.
func (p *scheduledtransactionRepo) CreateUserTransactionTx(req model_scheduled_transactions.ScheduledTransaction, tx *sql.Tx) (model_scheduled_transactions.ScheduledTransaction, *sql.Tx, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into schedule_transactions (type, status, user_id, reference, amount,paid_period) values ($1, $2, $3, $4, $5, $6)
	returning id, type, status, user_id, reference, amount, paid_period, created_at, updated_at`
	var model model_scheduled_transactions.ScheduledTransaction
	err := tx.QueryRowContext(ctx, stmt, req.Type, req.Status, req.UserID, req.Reference, req.Amount, req.PaidPeriod).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.PaidPeriod, &model.CreatedAt, &model.UpdatedAt)
	return model, tx, err
}

// stmt := `insert into accounts (name, number, user_id, recipient_code, bank_code)
// values ($1, $2, $3, $4, $5) returning id, name, number, user_id, created_at, updated_at
// Type      string    `json:"type"`
// 	Status    string    `json:"status"`
// 	UserID    uuid.UUID `json:"user_id"`
// 	Reference string    `json:"reference"`
// 	Amount    int       `json:"amount"
