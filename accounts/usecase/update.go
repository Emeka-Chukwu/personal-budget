package usecase_account

import (
	model_account "personal-budget/accounts/model"
	"personal-budget/payment"

	"github.com/google/uuid"
)

// Update implements AccountUsecase.
func (us *accountUsecase) Update(model model_account.Account, id uuid.UUID) (*model_account.Account, error) {
	account, err := us.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if account.BankCode != model.BankCode {
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
		account.RecipientCode = data.Data.RecipientCode
	}

	account.Name = model.Name
	account.Number = model.Number

	return us.repo.Update(*account, account.ID)
}
