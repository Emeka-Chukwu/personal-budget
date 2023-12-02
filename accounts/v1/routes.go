package account_v1

import (
	model_account "personal-budget/accounts/model"
	usecase_account "personal-budget/accounts/usecase"
	"personal-budget/middleware"

	"github.com/gin-gonic/gin"
)

func NewAccountsRoutes(router *gin.RouterGroup, usecase usecase_account.AccountUsecase) {
	acctHandler := NewAccountHandler(usecase)
	route := router.Group("/accounts")
	route.POST("/create", middleware.ValidatorMiddleware[model_account.Account], acctHandler.CreateAccount)
	route.PUT("/update", middleware.ValidatorMiddleware[model_account.Account], acctHandler.Update)
	route.GET("/:id/fetch", acctHandler.FetchAccount)
	route.GET("/list", acctHandler.List)
	route.DELETE("/:id/delete", acctHandler.Delete)
}
