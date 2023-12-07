package wallet_v1

import (
	"fmt"
	"net/http"
	usecase_account "personal-budget/accounts/usecase"
	"personal-budget/payment"
	"personal-budget/shared"
	model_transaction "personal-budget/transactions/model"
	usecases_transaction "personal-budget/transactions/usecases"
	usecase_user "personal-budget/users/usecase"
	"personal-budget/util"
	wallet_model "personal-budget/wallet/model"
	usecase_wallet "personal-budget/wallet/usecase"

	"github.com/gin-gonic/gin"
)

type walletHandler struct {
	transusecase  usecases_transaction.TransactionUsecase
	walletusecase usecase_wallet.WalletUsecase
	payService    payment.PaymentInterface
	accountUsecae usecase_account.AccountUsecase
	userUsecase   usecase_user.UsecaseUser
}

// Fetch implements WalletHandler.
func (handler *walletHandler) Fetch(ctx *gin.Context) {
	payload := shared.GetAuthsPayload(ctx)
	wallet, err := handler.walletusecase.Fetch(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": wallet})
}

// InitiateFundWallet implements WalletHandler.
func (handler *walletHandler) InitiateFundWallet(ctx *gin.Context) {
	request := util.GetBody[payment.PayloadInit](ctx)
	resp, err := handler.payService.InitializePayment(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// Transfer implements WalletHandler.
func (handler *walletHandler) Transfer(ctx *gin.Context) {
	request := util.GetBody[wallet_model.WithdrawalRequest](ctx)
	payload := shared.GetAuthsPayload(ctx)
	_, err := handler.userUsecase.GetUserByEmail(request.ReceiveEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = handler.walletusecase.Transfer(payload.UserId, request.ReceiveEmail, request.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	touser, err := handler.userUsecase.GetUserByEmail(request.ReceiveEmail)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionParamsFrom := model_transaction.Transaction{Type: "debit", Status: "success",
		Reference: fmt.Sprintf("%s", util.RandomString(16)), UserID: payload.UserId, Amount: request.Amount,
	}
	transactionParamsTo := model_transaction.Transaction{Type: "credit", Status: "success",
		Reference: fmt.Sprintf("%s", util.RandomString(16)), UserID: touser.ID, Amount: request.Amount,
	}
	handler.transusecase.CreateUserTransaction(transactionParamsFrom)
	handler.transusecase.CreateUserTransaction(transactionParamsTo)
	ctx.JSON(http.StatusOK, gin.H{"data": "Transfer was successful"})

}

// Withdrawal implements WalletHandler.
func (handler *walletHandler) Withdrawal(ctx *gin.Context) {
	request := util.GetBody[payment.InitiateTransfer](ctx)
	payload := shared.GetAuthsPayload(ctx)
	fromAccount, err := handler.accountUsecae.Get(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	request.Recipient = fromAccount.RecipientCode
	resp, err := handler.payService.Create(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	transactionParams := model_transaction.Transaction{Type: "credit", Status: "success",
		Reference: resp.Data.Reference, UserID: payload.UserId, Amount: int(request.Amount),
	}
	handler.transusecase.CreateUserTransaction(transactionParams)
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

type WalletHandler interface {
	InitiateFundWallet(ctx *gin.Context)
	Withdrawal(ctx *gin.Context)
	Fetch(ctx *gin.Context)
	Transfer(ctx *gin.Context)
}

func NewWalletHandler(trans usecases_transaction.TransactionUsecase, walletusecase usecase_wallet.WalletUsecase) WalletHandler {
	return &walletHandler{transusecase: trans, walletusecase: walletusecase}
}
