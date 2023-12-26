package schedule_payment_usecase

import (
	"errors"
	repositories_account "personal-budget/accounts/repositories"
	schedule_payment_model "personal-budget/schedule/model"
	schedule_payment_respositories "personal-budget/schedule/repositories"
	repositories_wallet "personal-budget/wallet/repositories"

	"github.com/google/uuid"
)

// WalletRepo

type ScheduledPaymentsUsecase interface {
	CreatePlan(model schedule_payment_model.SchedulePayment) (schedule_payment_model.SchedulePayment, error)
	FetchPlan(id uuid.UUID) (schedule_payment_model.SchedulePaymentPlan, error)
	ListPlan(userId uuid.UUID) ([]schedule_payment_model.SchedulePaymentPlan, error)
	CreatePlanTx(req schedule_payment_model.SchedulePayment) (schedule_payment_model.SchedulePayment, error)
}

type scheduledPaymentsUsecase struct {
	repo        schedule_payment_respositories.SchedulePaymentRepositories
	walletsRepo repositories_wallet.WalletRepo
	accountRepo repositories_account.AccountRepository
}

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

// FetchPlan implements ScheduledPaymentsUsecase.
func (p *scheduledPaymentsUsecase) FetchPlan(id uuid.UUID) (schedule_payment_model.SchedulePaymentPlan, error) {
	return p.repo.FetchPlan(id)
}

// ListPlan implements ScheduledPaymentsUsecase.
func (p *scheduledPaymentsUsecase) ListPlan(userId uuid.UUID) ([]schedule_payment_model.SchedulePaymentPlan, error) {
	return p.repo.ListPlan(userId)
}

func NewScheduledPaymentsUsecase(repo schedule_payment_respositories.SchedulePaymentRepositories,
	walletRepo repositories_wallet.WalletRepo, accountRepo repositories_account.AccountRepository) ScheduledPaymentsUsecase {
	return &scheduledPaymentsUsecase{repo: repo, walletsRepo: walletRepo, accountRepo: accountRepo}
}
