package webhook_v1

import (
	"net/http"
	webhook_usecase "personal-budget/webhook/usecase"

	"github.com/gin-gonic/gin"
)

type webhookHandler struct {
	usecase webhook_usecase.Webhookusecase
}

// PaystackWebhook implements WebhookHandler.
func (handler *webhookHandler) PaystackWebhook(ctx *gin.Context) {
	data, err := handler.usecase.PayStackWebhook(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

type WebhookHandler interface {
	PaystackWebhook(ctx *gin.Context)
}

func NewWebhookHandler(usecase webhook_usecase.Webhookusecase) WebhookHandler {
	return &webhookHandler{usecase: usecase}
}
