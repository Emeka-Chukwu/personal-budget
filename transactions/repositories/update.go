package repositories_transaction

import model_transaction "personal-budget/transactions/model"

// UpdateUserTransaction implements TransactionRepo.
func (*transactionRepo) UpdateUserTransaction(reference string, status string) (model_transaction.Transaction, error) {
	panic("unimplemented")
}
