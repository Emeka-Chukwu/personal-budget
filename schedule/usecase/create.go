package schedule_payment_usecase

import (
	"errors"
	schedule_payment_model "personal-budget/schedule/model"
	"personal-budget/shared"
	"time"
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
	noOfDays := int(req.Duration / req.Periods)
	noOfHours := req.Duration % req.Periods
	payDate := time.Now().AddDate(0, 0, noOfDays)
	payDate = payDate.Add(time.Hour * time.Duration(noOfHours))
	req.PayDate = payDate
	wallet, err := p.walletsRepo.Fetch(req.UserId)
	if err != nil {
		return schedule_payment_model.SchedulePayment{}, err
	}
	if req.Amount+shared.ScheduledPaymentCharge > wallet.Balance {
		return schedule_payment_model.SchedulePayment{}, errors.New("insufficient balance")
	}
	tx, err := p.walletsRepo.DebitFromWalletTx(req.UserId, req.Amount+shared.ScheduledPaymentCharge)
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
