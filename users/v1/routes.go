package users_v1

import (
	"personal-budget/middleware"
	model_user "personal-budget/users/models"
	usecase_user "personal-budget/users/usecase"

	"github.com/gin-gonic/gin"
)

func NewUserRoutes(router *gin.RouterGroup, usecase usecase_user.UsecaseUser) {
	userHandler := NewUserHandler(usecase)
	route := router.Group("/auths")
	route.POST("/register", middleware.ValidatorMiddleware[model_user.User], userHandler.RegisterUser)
	route.POST("/login", middleware.ValidatorMiddleware[model_user.UserLogin], userHandler.LoginUser)
}
