package usecase_wallet

import (
	"personal-budget/payment"

	"github.com/google/uuid"
)

// FundWallet implements WalletUsecase.
// / for payment web hook
func (repo *walletUsecase) InitiateFundWallet(payload payment.PayloadInit, userId uuid.UUID) (payment.Payload, error) {
	return repo.pay.InitializePayment(payload)

}
