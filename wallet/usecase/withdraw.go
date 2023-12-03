package usecase_wallet

import "github.com/google/uuid"

// Withdrawal implements WalletUsecase.
func (repo *walletUsecase) Withdrawal(userId uuid.UUID, amount int, callback func() error) error {
	panic("unimplemented")
}

// WithdrawalExample implements WalletUsecase.
func (repo *walletUsecase) WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error {
	return repo.walletRepo.WithdrawalExample(userId, amount, func(userId uuid.UUID, amount int) error {
		/// write o your paystack api call here
		return nil
	})
}
