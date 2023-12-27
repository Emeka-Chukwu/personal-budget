package repositories_scheduled_transactions

import (
	"database/sql"
	model_scheduled_transactions "personal-budget/schedule_transactions/model"

	"github.com/google/uuid"
)

type ScheduledTransactionRepo interface {
	GetUserTransaction(userId uuid.UUID) (model_scheduled_transactions.ScheduledTransaction, error)
	GetPlanScheduledTransaction(userId uuid.UUID) (model_scheduled_transactions.ScheduledTransaction, error)
	GetUserTransactions(userId uuid.UUID) ([]model_scheduled_transactions.ScheduledTransaction, error)
	CreateUserTransaction(model_scheduled_transactions.ScheduledTransaction) (model_scheduled_transactions.ScheduledTransaction, error)
	UpdateUserTransaction(reference, status string) (model_scheduled_transactions.ScheduledTransaction, error)
	GetUserTransactionByReference(reference string) (model_scheduled_transactions.ScheduledTransaction, error)
	UpdateUserTransactionType(reference string, typeD string) (model_scheduled_transactions.ScheduledTransaction, error)
	CreateUserTransactionTx(req model_scheduled_transactions.ScheduledTransaction, tx *sql.Tx) (model_scheduled_transactions.ScheduledTransaction, *sql.Tx, error)
}

type scheduledtransactionRepo struct {
	DB *sql.DB
}

// GetPlanScheduledTransaction implements ScheduledTransactionRepo.

func NewTransactionRepo(DB *sql.DB) ScheduledTransactionRepo {
	return &scheduledtransactionRepo{DB: DB}
}
