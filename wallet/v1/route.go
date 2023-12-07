package wallet_v1

import (
	usecase_account "personal-budget/accounts/usecase"
	"personal-budget/payment"
	usecases_transaction "personal-budget/transactions/usecases"
	usecase_user "personal-budget/users/usecase"
	usecase_wallet "personal-budget/wallet/usecase"

	"github.com/gin-gonic/gin"
)

func NewWebhooksRoutes(router *gin.RouterGroup,
	transusecase usecases_transaction.TransactionUsecase,
	walletusecase usecase_wallet.WalletUsecase,
	payService payment.PaymentInterface,
	accountUsecase usecase_account.AccountUsecase,
	userUsecase usecase_user.UsecaseUser) {
	walletHandler := NewWalletHandler(transusecase, walletusecase, accountUsecase, payService, userUsecase)
	route := router.Group("/wallets")
	route.GET("/balance", walletHandler.Fetch)
	route.POST("/fund-account", walletHandler.InitiateFundWallet)
	route.POST("/withdraw", walletHandler.Withdrawal)
	route.POST("/transfer", walletHandler.Transfer)
}
