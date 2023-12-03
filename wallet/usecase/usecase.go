package usecase_wallet

import (
	"personal-budget/token"
	repositories_users "personal-budget/users/repositories"
	wallet_model "personal-budget/wallet/model"
	repositories_wallet "personal-budget/wallet/repositories"

	"github.com/google/uuid"
)

type WalletUsecase interface {
	FundWallet(amount int, userId uuid.UUID) (wallet_model.Wallet, error)
	Withdrawal(userId uuid.UUID, amount int, callback func() error) error
	WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error
	Fetch(userId uuid.UUID) (wallet_model.Wallet, error)
	Transfer(fromUserId uuid.UUID, receiverEmail string, amount int) error
}

type walletUsecase struct {
	token      token.Maker
	walletRepo repositories_wallet.WalletRepo
	userRepo   repositories_users.UserAuthentication
}

func NewAccountUsecase(token token.Maker, repo repositories_wallet.WalletRepo,
	userRepo repositories_users.UserAuthentication) WalletUsecase {
	return &walletUsecase{walletRepo: repo, token: token, userRepo: userRepo}
}
