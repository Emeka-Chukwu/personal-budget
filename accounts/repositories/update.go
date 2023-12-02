package repositories_account

import (
	"context"
	model_account "personal-budget/accounts/model"
	"personal-budget/util"

	"github.com/google/uuid"
)

// Update implements AccountRepository.
func (repo *accountRepository) Update(model model_account.Account, id uuid.UUID) (*model_account.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update accounts set name=$1, number=$2 where id=$3 returning id, name, number, user_id, created_at, updated_at`
	var resp model_account.Account
	err := repo.DB.QueryRowContext(ctx, stmt, model.Name, model.Number, id).
		Scan(
			&resp.ID,
			&resp.Name,
			&resp.Number,
			&resp.UserId,
			&resp.CreatedAt,
			&resp.UpdatedAt,
		)
	return &resp, err
}
