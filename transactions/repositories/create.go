package repositories_transaction

import model_transaction "personal-budget/transactions/model"

// CreateUserTransaction implements TransactionRepo.
func (*transactionRepo) CreateUserTransaction(model model_transaction.Transaction) (model_transaction.Transaction, error) {
	panic("unimplemented")
}
