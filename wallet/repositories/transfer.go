package repositories_wallet

import (
	"github.com/google/uuid"
)

// Transfer implements WalletRepo.
func (repo *walletRepo) Transfer(fromUserId uuid.UUID, toUserId uuid.UUID, amount int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	stmt1 := `UPDATE wallets SET balance = balance - $1 WHERE user_id = $2`
	_, err = tx.Exec(stmt1, amount, fromUserId)
	if err != nil {
		tx.Rollback()
		return err
	}
	stmt2 := `UPDATE wallets SET balance = balance + $1 WHERE user_id = $2`
	_, err = tx.Exec(stmt2, amount, toUserId)
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
