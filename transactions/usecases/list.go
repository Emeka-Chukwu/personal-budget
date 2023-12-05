package usecases_transaction

import (
	model_transaction "personal-budget/transactions/model"

	"github.com/google/uuid"
)

// GetUserTransactions implements TransactionUsecase.
func (t *transactionUsecase) GetUserTransactions(userId uuid.UUID) ([]model_transaction.Transaction, error) {
	return t.repo.GetUserTransactions(userId)
}
