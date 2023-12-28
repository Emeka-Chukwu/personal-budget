package usecases_scheduled_transactions

import (
	model_scheduled_transactions "personal-budget/schedule_transactions/model"
	repositories_scheduled_transactions "personal-budget/schedule_transactions/repositories"

	"github.com/google/uuid"
)

type TransactionUsecase interface {
	GetUserTransactions(userId uuid.UUID) ([]model_scheduled_transactions.ScheduledTransaction, error)
	CreateUserTransaction(model model_scheduled_transactions.ScheduledTransaction) (model_scheduled_transactions.ScheduledTransaction, error)
}

type transactionUsecase struct {
	repo repositories_scheduled_transactions.ScheduledTransactionRepo
}

// CreateUserTransaction implements TransactionUsecase.
func (tran *transactionUsecase) CreateUserTransaction(model model_scheduled_transactions.ScheduledTransaction) (model_scheduled_transactions.ScheduledTransaction, error) {
	return tran.repo.CreateUserTransaction(model)
	// CreateUserTransaction(model)
}

func NewTransactionUsecase(repo repositories_scheduled_transactions.ScheduledTransactionRepo) TransactionUsecase {
	return &transactionUsecase{repo: repo}
}
