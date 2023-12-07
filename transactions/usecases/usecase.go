package usecases_transaction

import (
	model_transaction "personal-budget/transactions/model"
	repositories_transaction "personal-budget/transactions/repositories"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	GetUserTransactions(userId uuid.UUID) ([]model_transaction.Transaction, error)
	CreateUserTransaction(model model_transaction.Transaction) (model_transaction.Transaction, error)
}

type transactionUsecase struct {
	repo repositories_transaction.TransactionRepo
}

// CreateUserTransaction implements TransactionUsecase.
func (tran *transactionUsecase) CreateUserTransaction(model model_transaction.Transaction) (model_transaction.Transaction, error) {
	return tran.repo.CreateUserTransaction(model)
}

func NewTransactionUsecase(repo repositories_transaction.TransactionRepo) TransactionUsecase {
	return &transactionUsecase{repo: repo}
}
