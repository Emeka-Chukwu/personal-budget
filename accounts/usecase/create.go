package usecase_account

import (
	model_account "personal-budget/accounts/model"
	"personal-budget/payment"
)

// Create implements AccountUsecase.
func (us *accountUsecase) Create(model model_account.Account) (*model_account.Account, error) {
	// accountValidation := util.NewApiCallInterface(us.config)
	// accountValidation.PaystackApiCall(model.Number, model.BankCode)
	recipient := payment.CreateRecipient{
		Type:          "nuban",
		Name:          model.Name,
		AccountNumber: model.Number,
		BankCode:      model.BankCode,
		Currency:      "NGN",
	}
	data, err := us.pay.GetRecipientCode(recipient)
	if err != nil {
		return nil, err
	}
	model.RecipientCode = data.Data.RecipientCode
	return us.repo.Create(model)
}
