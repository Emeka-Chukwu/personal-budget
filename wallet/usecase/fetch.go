package usecase_wallet

import (
	wallet_model "personal-budget/wallet/model"

	"github.com/google/uuid"
)

// Fetch implements WalletUsecase.
func (repo *walletUsecase) Fetch(userId uuid.UUID) (wallet_model.Wallet, error) {
	return repo.walletRepo.Fetch(userId)

}
