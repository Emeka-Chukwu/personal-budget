package schedule_payment_usecase

import (
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
