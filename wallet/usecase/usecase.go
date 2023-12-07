package usecase_wallet

import (
	"personal-budget/payment"
	"personal-budget/token"
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
	token      token.Maker
	walletRepo repositories_wallet.WalletRepo
	userRepo   repositories_users.UserAuthentication
	pay        payment.PaymentInterface
	config     util.Config
}

// WithdrawalExample implements WalletUsecase.
func (*walletUsecase) WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error {
	panic("unimplemented")
}

func NewAccountUsecase(token token.Maker, repo repositories_wallet.WalletRepo,
	userRepo repositories_users.UserAuthentication, config util.Config) WalletUsecase {
	return &walletUsecase{walletRepo: repo, token: token, userRepo: userRepo, config: config}
}
