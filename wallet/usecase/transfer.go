package usecase_wallet

import (
	"errors"

	"github.com/google/uuid"
)

// Transfer implements WalletUsecase.
func (repo *walletUsecase) Transfer(fromUserId uuid.UUID, receiverEmail string, amount int) error {
	fromWallet, err := repo.walletRepo.Fetch(fromUserId)
	if err != nil {
		return err
	}
	if amount > fromWallet.Balance {
		return errors.New("insufficient amount, please top up")
	}
	toUserAcct, err := repo.userRepo.GetUserByEmail(receiverEmail)
	if err != nil {
		return err
	}
	err = repo.walletRepo.Transfer(fromUserId, toUserAcct.ID, amount)
	if err != nil {
		return err
	}

	return nil

}
