package transaction_v1

import (
	usecases_transaction "personal-budget/transactions/usecases"

	"github.com/gin-gonic/gin"
)

func NewTransactionRoutes(router *gin.RouterGroup, usecase usecases_transaction.TransactionUsecase) {
	transHandler := NewTransactionHandler(usecase)
	route := router.Group("/transactions")
	route.GET("/list", transHandler.GetUserTransactions)

}
