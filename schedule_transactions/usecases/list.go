package usecases_scheduled_transactions

import (
	model_scheduled_transactions "personal-budget/schedule_transactions/model"

	"github.com/google/uuid"
)

// GetUserTransactions implements TransactionUsecase.
func (t *transactionUsecase) GetUserTransactions(userId uuid.UUID) ([]model_scheduled_transactions.ScheduledTransaction, error) {
	return t.repo.GetUserTransactions(userId)
}
