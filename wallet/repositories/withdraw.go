package repositories_wallet

import (
	"github.com/google/uuid"
)

// Withdrawal implements WalletRepo.
func (repo *walletRepo) Withdrawal(userId uuid.UUID, amount int, callback func() error) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	stmt := `update wallets set balance = balance - $1 where user_id=$2`
	_, err = tx.Exec(stmt, amount, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = callback()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil

}

// WithdrawalExampl implements WalletRepo.
func (repo *walletRepo) WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	stmt := `update wallets set balance = balance - $1 where user_id=$2`
	_, err = tx.Exec(stmt, amount, userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = callback(userId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
