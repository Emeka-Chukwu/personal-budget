package schedule_payment_respositories

import (
	schedule_payment_model "personal-budget/schedule/model"

	"github.com/google/uuid"
)

// ListPlan implements SchedulePaymentRepositories.
func (*schedulePaymentRepositories) ListPlan(userId uuid.UUID) ([]schedule_payment_model.SchedulePaymentPlan, error) {
	panic("unimplemented")
}
