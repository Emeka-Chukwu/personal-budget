package webhook_v1

import (
	webhook_usecase "personal-budget/webhook/usecase"

	"github.com/gin-gonic/gin"
)

func NewWebhooksRoutes(router *gin.RouterGroup, usecase webhook_usecase.Webhookusecase) {
	webhookHandler := NewWebhookHandler(usecase)
	route := router.Group("/webhooks")
	route.POST("/paystack", webhookHandler.PaystackWebhook)

}
