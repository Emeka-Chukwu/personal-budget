package users_v1

import (
	"net/http"
	model_user "personal-budget/users/models"
	usecase_user "personal-budget/users/usecase"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type UserHandler interface {
	LoginUser(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
}

type userhandler struct {
	usecase usecase_user.UsecaseUser
}

// LoginUser implements UserHandler.
func (*userhandler) LoginUser(ctx *gin.Context) {
	panic("unimplemented")
}

// RegisterUser implements UserHandler.
func (handler *userhandler) RegisterUser(ctx *gin.Context) {
	req := util.GetBody[model_user.User](ctx)
	resp, err := handler.usecase.Register(req)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, gin.H{"error": "email already exist"})
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"data": resp})
}

func NewUserHandler(usecase usecase_user.UsecaseUser) UserHandler {
	return &userhandler{usecase: usecase}
}
