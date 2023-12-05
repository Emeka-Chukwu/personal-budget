package usecases_transaction

import (
	model_transaction "personal-budget/transactions/model"
	repositories_transaction "personal-budget/transactions/repositories"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	GetUserTransactions(userId uuid.UUID) ([]model_transaction.Transaction, error)
}

type transactionUsecase struct {
	repo repositories_transaction.TransactionRepo
}

func NewTransactionUsecase(repo repositories_transaction.TransactionRepo) TransactionUsecase {
	return &transactionUsecase{repo: repo}
}
