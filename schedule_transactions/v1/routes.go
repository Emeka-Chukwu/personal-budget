package transaction_scheduled_v1

import (
	usecases_scheduled_transactions "personal-budget/schedule_transactions/usecases"

	"github.com/gin-gonic/gin"
)

func NewTransactionRoutes(router *gin.RouterGroup, usecase usecases_scheduled_transactions.TransactionUsecase) {
	transHandler := NewTransactionHandler(usecase)
	route := router.Group("/scheduled-transactions")
	route.GET("/list", transHandler.GetUserTransactions)

}
