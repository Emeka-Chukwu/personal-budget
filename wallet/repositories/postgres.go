package repositories_wallet

import (
	"database/sql"
	wallet_model "personal-budget/wallet/model"

	"github.com/google/uuid"
)

type WalletRepo interface {
	FundAccount(amount int, userId uuid.UUID) (wallet_model.Wallet, error)
	Withdrawal(userId uuid.UUID, amount int, callback func() error) error
	WithdrawalExample(userId uuid.UUID, amount int, callback func(userId uuid.UUID, amount int) error) error
	Fetch(userId uuid.UUID) (wallet_model.Wallet, error)
	Transfer(fromUserId, toUserId uuid.UUID, amount int) error
}

type walletRepo struct {
	DB *sql.DB
}

func NewWalletRepo(db *sql.DB) WalletRepo {
	return &walletRepo{DB: db}
}
