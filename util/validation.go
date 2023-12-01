package util

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func ValidateInput(payload interface{}) (map[string]string, error) {
	v := validator.New()
	err := v.Struct(payload)

	errors := make(map[string]string)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {

			errors[strings.ToLower(e.Field())] = e.ActualTag()
		}
	}
	return errors, err
}

func GetBody[T any](c *gin.Context) T {
	var payload T
	err := c.ShouldBindBodyWith(&payload, binding.JSON)
	if err != nil {
		return payload
	}
	return payload
}
