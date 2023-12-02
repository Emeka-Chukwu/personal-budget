package constant_account

import (
	"personal-budget/token"

	"github.com/gin-gonic/gin"
)

var (
	authorizationPayloadKey = "authorization_payload"
)

func GetAuthsPayload(ctx *gin.Context) *token.Payload {
	pay, _ := ctx.Get(authorizationPayloadKey)
	return pay.(*token.Payload)
}
