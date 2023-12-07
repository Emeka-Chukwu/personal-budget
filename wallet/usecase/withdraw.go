package usecase_wallet

import (
	"personal-budget/payment"

	"github.com/google/uuid"
)

// Withdrawal implements WalletUsecase.
func (repo *walletUsecase) Withdrawal(userId uuid.UUID, amount int, request payment.InitiateTransfer, callback func(transferRequst payment.InitiateTransfer) (payment.TransferResponse, error)) (payment.TransferResponse, error) {
	data, err := repo.walletRepo.Withdrawal(userId, amount, request, func(transferRequst payment.InitiateTransfer) (payment.TransferResponse, error) {
		return repo.pay.Create(request)
	})
	return data, err
}
