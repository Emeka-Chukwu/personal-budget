package usecase_account

import (
	model_account "personal-budget/accounts/model"

	"github.com/google/uuid"
)

// Get implements AccountUsecase.
func (us *accountUsecase) Get(id uuid.UUID) (*model_account.Account, error) {
	return us.repo.Get(id)
}

// List implements AccountUsecase.
func (us *accountUsecase) List(id uuid.UUID) ([]model_account.Account, error) {
	return us.repo.List(id)
}

// GetByID implements AccountUsecase.
func (us *accountUsecase) GetByID(id uuid.UUID) (*model_account.Account, error) {
	return us.repo.GetByID(id)
}
