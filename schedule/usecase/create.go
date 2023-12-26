package schedule_payment_usecase

import (
	"errors"
	schedule_payment_model "personal-budget/schedule/model"
)

// CreatePlanTx implements ScheduledPaymentsUsecase.
func (p *scheduledPaymentsUsecase) CreatePlanTx(req schedule_payment_model.SchedulePayment) (schedule_payment_model.SchedulePayment, error) {
	account, err := p.accountRepo.GetByID(req.AccountId)
	if err != nil {
		return schedule_payment_model.SchedulePayment{}, err
	}
	if account.UserId != req.UserId {
		return schedule_payment_model.SchedulePayment{}, errors.New("access denied")
	}
	tx, err := p.walletsRepo.DebitFromWalletTx(req.UserId, req.Amount)
	if err != nil {
		tx.Rollback()
		return schedule_payment_model.SchedulePayment{}, err
	}
	data, txns, err := p.repo.CreatePlanTx(req, tx)
	if err != nil {
		tx.Rollback()
		return schedule_payment_model.SchedulePayment{}, err
	}
	if err = txns.Commit(); err != nil {
		tx.Rollback()
		return schedule_payment_model.SchedulePayment{}, err
	}
	return data, nil
}

// CreatePlan implements ScheduledPaymentsUsecase.
func (*scheduledPaymentsUsecase) CreatePlan(model schedule_payment_model.SchedulePayment) (schedule_payment_model.SchedulePayment, error) {
	panic("unimplemented")
}
