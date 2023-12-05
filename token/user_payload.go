package token

import "github.com/gin-gonic/gin"

var (
	authorizationPayloadKey = "authorization_payload"
)

func GetAuthsPayload(ctx *gin.Context) *Payload {
	pay, _ := ctx.Get(authorizationPayloadKey)
	return pay.(*Payload)
}
