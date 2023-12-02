package usecase_account

import (
	model_account "personal-budget/accounts/model"

	"github.com/google/uuid"
)

// Update implements AccountUsecase.
func (us *accountUsecase) Update(model model_account.Account, id uuid.UUID) (*model_account.Account, error) {
	account, err := us.repo.Get(id)
	if err != nil {
		return nil, err
	}
	account.Name = model.Name
	account.Number = model.Number
	return us.repo.Update(*account, account.ID)
}
