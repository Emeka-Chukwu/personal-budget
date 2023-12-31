package usecase_wallet

import (
	"personal-budget/payment"
	repositories_users "personal-budget/users/repositories"
	"personal-budget/util"
	wallet_model "personal-budget/wallet/model"
	repositories_wallet "personal-budget/wallet/repositories"

	"github.com/google/uuid"
)

type WalletUsecase interface {
	InitiateFundWallet(payload payment.PayloadInit, userId uuid.UUID) (payment.Payload, error)
	Withdrawal(userId uuid.UUID, amount int, request payment.InitiateTransfer, callback func(transferRequst payment.InitiateTransfer) (payment.TransferResponse, error)) (payment.TransferResponse, error)
	WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error
	Fetch(userId uuid.UUID) (wallet_model.Wallet, error)
	Transfer(fromUserId uuid.UUID, receiverEmail string, amount int) error
}

type walletUsecase struct {
	walletRepo repositories_wallet.WalletRepo
	userRepo   repositories_users.UserAuthentication
	pay        payment.PaymentInterface
	config     util.Config
}

// WithdrawalExample implements WalletUsecase.
func (*walletUsecase) WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error {
	panic("unimplemented")
}

func NewWalletUsecase(repo repositories_wallet.WalletRepo,
	userRepo repositories_users.UserAuthentication, config util.Config, pay payment.PaymentInterface) WalletUsecase {
	return &walletUsecase{walletRepo: repo, userRepo: userRepo, config: config, pay: pay}
}
