package usecase_wallet

import (
	wallet_model "personal-budget/wallet/model"

	"github.com/google/uuid"
)

// FundWallet implements WalletUsecase.
// / for payment web hook
func (repo *walletUsecase) FundWallet(amount int, userId uuid.UUID) (wallet_model.Wallet, error) {
	// panic("unimplemented")
	///call the card payment method here and thats all remember you'll be using webhook for others will entirely be a n
	return repo.walletRepo.FundAccount(amount, userId)
}
