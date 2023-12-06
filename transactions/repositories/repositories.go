package repositories_transaction

import (
	"database/sql"
	model_transaction "personal-budget/transactions/model"

	"github.com/google/uuid"
)

// func (p *paymentInterface) GetUserTransaction(requ Transaction) (Transaction, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
// 	defer cancel()
// 	stmt := `select id, type, status, user_id, reference, amount, created_at, updated_at from transactions where user_id=$1 limit 1`
// 	var resp Transaction
// 	err := p.DB.QueryRowContext(ctx, stmt, id).
// 		Scan(&model.ID, &model.Name, &model.Number, &model.UserId, &model.RecipientCode, &model.BankCode, &model.CreatedAt, &model.UpdatedAt)

// 	return &model, err
// }

// func (p *paymentInterface) GetUserTransaction(requ Transaction) (Transaction, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
// 	defer cancel()
// 	stmt := `select id, type, status, user_id, reference, amount, created_at, updated_at from transactions where user_id=$1 limit 1`
// 	var resp Transaction
// 	err := p.DB.QueryRowContext(ctx, stmt, id).
// 		Scan(&model.ID, &model.Name, &model.Number, &model.UserId, &model.RecipientCode, &model.BankCode, &model.CreatedAt, &model.UpdatedAt)

// 	return &model, err
// }

type TransactionRepo interface {
	GetUserTransaction(userId uuid.UUID) (model_transaction.Transaction, error)
	GetUserTransactions(userId uuid.UUID) ([]model_transaction.Transaction, error)
	CreateUserTransaction(model model_transaction.Transaction) (model_transaction.Transaction, error)
	UpdateUserTransaction(reference, status string) (model_transaction.Transaction, error)
	GetUserTransactionByReference(reference string) (model_transaction.Transaction, error)
	UpdateUserTransactionType(reference string, typeD string) (model_transaction.Transaction, error)
}

type transactionRepo struct {
	DB *sql.DB
}

func NewTransactionRepo(DB *sql.DB) TransactionRepo {
	return &transactionRepo{DB: DB}
}
