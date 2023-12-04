package repositories_transaction

import (
	model_transaction "personal-budget/transactions/model"

	"github.com/google/uuid"
)

// GetUserTransaction implements TransactionRepo.
func (*transactionRepo) GetUserTransaction(userId uuid.UUID) (model_transaction.Transaction, error) {
	panic("unimplemented")
}

// GetUserTransactionByReference implements TransactionRepo.
func (*transactionRepo) GetUserTransactionByReference(reference string) (model_transaction.Transaction, error) {
	panic("unimplemented")
}

// GetUserTransactions implements TransactionRepo.
func (*transactionRepo) GetUserTransactions(userId uuid.UUID) ([]model_transaction.Transaction, error) {
	panic("unimplemented")
}
