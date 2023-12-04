package repositories_account

import (
	"context"
	model_account "personal-budget/accounts/model"
	"personal-budget/util"
)

// Create implements AccountRepository.
func (repo *accountRepository) Create(model model_account.Account) (*model_account.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into accounts (name, number, user_id, recipient_code, bank_code)
		values ($1, $2, $3, $4, $5) returning id, name, number, user_id, created_at, updated_at`
	var resp model_account.Account
	err := repo.DB.QueryRowContext(ctx, stmt,
		model.Name,
		model.Number,
		model.UserId,
		model.RecipientCode,
		model.BankCode,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Number,
		&resp.UserId,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	return &resp, err
}
