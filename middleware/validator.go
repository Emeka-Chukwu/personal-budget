package middleware

import (
	"net/http"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidatorMiddleware[T any](c *gin.Context) {
	var payload T
	err := c.ShouldBindBodyWith(&payload, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	errs, err := util.ValidateInput(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		c.Abort()
		return
	}
	c.Set("body", payload)

	c.Next()
}
