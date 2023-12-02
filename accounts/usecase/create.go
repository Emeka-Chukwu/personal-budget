package usecase_account

import (
	model_account "personal-budget/accounts/model"
	"personal-budget/util"
)

// Create implements AccountUsecase.
func (us *accountUsecase) Create(model model_account.Account) (*model_account.Account, error) {
	accountValidation := util.NewApiCallInterface(us.config)
	err := accountValidation.PaystackApiCall(model.Number, model.BankCode)
	if err != nil {
		return nil, err
	}
	return us.repo.Create(model)
}
