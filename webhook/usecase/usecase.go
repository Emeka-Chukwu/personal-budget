package webhook_usecase

import (
	"personal-budget/payment"
	repositories_transaction "personal-budget/transactions/repositories"
	repositories_wallet "personal-budget/wallet/repositories"

	"github.com/gin-gonic/gin"
)

type Webhookusecase interface {
	WithdrawalWebhook(payment.DataTransfer) (payment.TransferResponse, error)
	FundWebhook(payment.DataTransfer) (payment.TransferResponse, error)
	PayStackWebhook(ctx *gin.Context) (any, error)
	IsFromPaystackSource(ctx *gin.Context) (bool, error)
}
type webhookusecase struct {
	walletRepo      repositories_wallet.WalletRepo
	transactionRepo repositories_transaction.TransactionRepo
}

// FundWebhook implements Webhookusecase.
func (*webhookusecase) FundWebhook(payment.DataTransfer) (payment.TransferResponse, error) {
	panic("unimplemented")
}

// IsFromPaystackSource implements Webhookusecase.
func (*webhookusecase) IsFromPaystackSource(ctx *gin.Context) (bool, error) {
	panic("unimplemented")
}

// PayStackWebhook implements Webhookusecase.
func (*webhookusecase) PayStackWebhook(ctx *gin.Context) (any, error) {
	panic("unimplemented")
}

// WithdrawalWebhook implements Webhookusecase.
func (*webhookusecase) WithdrawalWebhook(payment.DataTransfer) (payment.TransferResponse, error) {
	panic("unimplemented")
}

func NewWebhookusecase(walletRepo repositories_wallet.WalletRepo,
	transactionRepo repositories_transaction.TransactionRepo) Webhookusecase {
	return &webhookusecase{walletRepo: walletRepo, transactionRepo: transactionRepo}
}
