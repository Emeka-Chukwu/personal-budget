package account_v1

import (
	"net/http"
	constant_account "personal-budget/accounts/constant"
	model_account "personal-budget/accounts/model"
	usecase_account "personal-budget/accounts/usecase"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type AccountHandler interface {
	CreateAccount(ctx *gin.Context)
	FetchAccount(ctx *gin.Context)
	List(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type accounthandler struct {
	usecase usecase_account.AccountUsecase
}

// CreateAccount implements AccountHandler.
func (handler *accounthandler) CreateAccount(ctx *gin.Context) {
	req := util.GetBody[model_account.Account](ctx)
	resp, err := handler.usecase.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": resp})
}

// Delete implements AccountHandler.
func (handler *accounthandler) Delete(ctx *gin.Context) {
	req := util.GetUrlParams[model_account.AccountParam](ctx)
	err := handler.usecase.Delete(req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "record has been deleted"})
}

// FetchAccount implements AccountHandler.
func (handler *accounthandler) FetchAccount(ctx *gin.Context) {
	payload := constant_account.GetAuthsPayload(ctx)
	resp, err := handler.usecase.Get(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})

}

// List implements AccountHandler.
func (handler *accounthandler) List(ctx *gin.Context) {
	payload := constant_account.GetAuthsPayload(ctx)
	resp, err := handler.usecase.List(payload.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})

}

// Update implements AccountHandler.
func (handler *accounthandler) Update(ctx *gin.Context) {
	body := util.GetBody[model_account.Account](ctx)
	req := util.GetUrlParams[model_account.AccountParam](ctx)
	resp, err := handler.usecase.Update(body, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func NewAccountHandler(usecase usecase_account.AccountUsecase) AccountHandler {
	return &accounthandler{usecase: usecase}
}
