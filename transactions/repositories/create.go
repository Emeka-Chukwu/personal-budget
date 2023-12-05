package repositories_transaction

import (
	"context"
	model_transaction "personal-budget/transactions/model"
	"personal-budget/util"
)

// CreateUserTransaction implements TransactionRepo.
func (p *transactionRepo) CreateUserTransaction(req model_transaction.Transaction) (model_transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into transactions (type, status, user_id, reference, amount) values ($1, $2, $3, $4, $5)
	returning id, type, status, user_id, reference, amount, created_at, updated_at`
	var model model_transaction.Transaction
	err := p.DB.QueryRowContext(ctx, stmt, req.Type, req.Status, req.UserID, req.Reference, req.Amount).
		Scan(&model.ID, &model.Type, &model.Status, &model.UserID, &model.Reference, &model.Amount, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}

// stmt := `insert into accounts (name, number, user_id, recipient_code, bank_code)
// values ($1, $2, $3, $4, $5) returning id, name, number, user_id, created_at, updated_at
// Type      string    `json:"type"`
// 	Status    string    `json:"status"`
// 	UserID    uuid.UUID `json:"user_id"`
// 	Reference string    `json:"reference"`
// 	Amount    int       `json:"amount"
