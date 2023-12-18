package transaction_v1

import (
	"net/http"
	"personal-budget/token"
	usecases_transaction "personal-budget/transactions/usecases"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	usecase usecases_transaction.TransactionUsecase
}

// GetUserTransactions implements TransactionHandler.
func (handler *transactionHandler) GetUserTransactions(ctx *gin.Context) {
	payload := token.GetAuthsPayload(ctx)
	resp, err := handler.usecase.GetUserTransactions(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

type TransactionHandler interface {
	GetUserTransactions(ctx *gin.Context)
}

func NewTransactionHandler(usecase usecases_transaction.TransactionUsecase) TransactionHandler {
	return &transactionHandler{usecase: usecase}
}
