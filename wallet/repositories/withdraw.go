package repositories_wallet

import (
	"personal-budget/payment"

	"github.com/google/uuid"
)

// Withdrawal implements WalletRepo.
func (repo *walletRepo) Withdrawal(userId uuid.UUID, amount int, request payment.InitiateTransfer, callback func(transferRequst payment.InitiateTransfer) (payment.TransferResponse, error)) (payment.TransferResponse, error) {
	tx, err := repo.DB.Begin()
	if err != nil {
		return payment.TransferResponse{}, err
	}
	stmt := `update wallets set balance = balance - $1 where user_id=$2`
	_, err = tx.Exec(stmt, amount, userId)
	if err != nil {
		tx.Rollback()
		return payment.TransferResponse{}, err
	}
	resp, err := callback(request)
	if err != nil {
		tx.Rollback()
		return payment.TransferResponse{}, err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return payment.TransferResponse{}, err
	}
	return resp, nil

}
