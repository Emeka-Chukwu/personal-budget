package wallet_v1

import "github.com/gin-gonic/gin"

func NewWebhooksRoutes(router *gin.RouterGroup, usecase webhook_usecase.Webhookusecase) {
	webhookHandler := NewWebhookHandler(usecase)
	route := router.Group("/webhooks")
	route.POST("/paystack", webhookHandler.PaystackWebhook)

}
