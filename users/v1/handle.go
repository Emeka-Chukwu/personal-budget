package users_v1

import (
	"net/http"
	banks_data "personal-budget/bank/data"
	model_user "personal-budget/users/models"
	usecase_user "personal-budget/users/usecase"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type UserHandler interface {
	LoginUser(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
	FetchBanks(ctx *gin.Context)
	FetchBankById(ctx *gin.Context)
}

type userhandler struct {
	usecase usecase_user.UsecaseUser
}

// FetchBankById implements UserHandler.
func (*userhandler) FetchBankById(ctx *gin.Context) {
	params := util.GetUrlParams[BankParam](ctx)
	data := banks_data.FetchBankById(params.ID)
	if data == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// FetchBanks implements UserHandler.
func (*userhandler) FetchBanks(ctx *gin.Context) {
	data := banks_data.PaystackBanks
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

// LoginUser implements UserHandler.
func (handler *userhandler) LoginUser(ctx *gin.Context) {
	req := util.GetBody[model_user.UserLogin](ctx)
	req.ClientIP = ctx.ClientIP()
	req.UserAgent = ctx.GetHeader("User-Agent")
	resp, err := handler.usecase.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// RegisterUser implements UserHandler.
func (handler *userhandler) RegisterUser(ctx *gin.Context) {
	req := util.GetBody[model_user.User](ctx)
	req.ClientIP = ctx.ClientIP()
	req.UserAgent = ctx.GetHeader("User-Agent")
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
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": resp})
}

func NewUserHandler(usecase usecase_user.UsecaseUser) UserHandler {
	return &userhandler{usecase: usecase}
}

type BankParam struct {
	ID int64 `uri:"id" binding:"required"`
}
