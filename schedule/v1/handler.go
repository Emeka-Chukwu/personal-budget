package schedule_payment_v1

import (
	"net/http"
	schedule_payment_model "personal-budget/schedule/model"
	schedule_payment_usecase "personal-budget/schedule/usecase"
	"personal-budget/shared"
	"personal-budget/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ScheduledPaymentHandler interface {
	CreatePlan(*gin.Context)
	FetchPlan(*gin.Context)
	ListPlan(*gin.Context)
	CreatePlanTx(*gin.Context)
}

type scheduledPaymentHandler struct {
	usecase schedule_payment_usecase.ScheduledPaymentsUsecase
}

// CreatePlan implements ScheduledPaymentHandler.
func (p *scheduledPaymentHandler) CreatePlan(c *gin.Context) {
	panic("unimplemented")

}

// CreatePlanTx implements ScheduledPaymentHandler.
func (p *scheduledPaymentHandler) CreatePlanTx(c *gin.Context) {
	payload := util.GetBody[schedule_payment_model.SchedulePayment](c)
	authPayload := shared.GetAuthsPayload(c)
	payload.UserId = authPayload.UserId
	resp, err := p.usecase.CreatePlanTx(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

// FetchPlan implements ScheduledPaymentHandler.
func (p *scheduledPaymentHandler) FetchPlan(c *gin.Context) {
	params := util.GetUrlParams[ParamID](c)
	paramsId, _ := uuid.Parse(params.ID)
	resp, err := p.usecase.FetchPlan(paramsId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

// ListPlan implements ScheduledPaymentHandler.
func (p *scheduledPaymentHandler) ListPlan(c *gin.Context) {
	authPayload := shared.GetAuthsPayload(c)
	userID := authPayload.UserId
	resp, err := p.usecase.ListPlan(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

func NewScheduledPaymentHandler(usecase schedule_payment_usecase.ScheduledPaymentsUsecase) ScheduledPaymentHandler {
	return &scheduledPaymentHandler{usecase: usecase}
}

type ParamID struct {
	ID string `uri:"id" binding:"required"`
}
