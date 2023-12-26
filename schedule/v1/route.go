package schedule_payment_v1

import (
	"personal-budget/middleware"
	schedule_payment_model "personal-budget/schedule/model"
	schedule_payment_usecase "personal-budget/schedule/usecase"

	"github.com/gin-gonic/gin"
)

func NewScheduledPaymentRoutes(router *gin.RouterGroup, usecase schedule_payment_usecase.ScheduledPaymentsUsecase) {
	planHandler := NewScheduledPaymentHandler(usecase)
	route := router.Group("/plans")
	route.POST("/create", middleware.ValidatorMiddleware[schedule_payment_model.SchedulePayment], planHandler.CreatePlanTx)
	route.GET("/fetch/:id", planHandler.FetchPlan)
	route.GET("/fetch/list", planHandler.ListPlan)
}
