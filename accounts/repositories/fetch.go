package repositories_account

import (
	"context"
	model_account "personal-budget/accounts/model"
	"personal-budget/util"

	"github.com/google/uuid"
)

// Get implements AccountRepository.
func (repo *accountRepository) Get(id uuid.UUID) (*model_account.Account, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	// defer cancel()
	// stmp := `select id, name,  created_at, updated_at from labels where id=$1`
	// var label resp.LabelResponse
	// err := lab.DB.QueryRowContext(ctx, stmp, id).
	// 	Scan(
	// 		&label.ID,
	// 		&label.Name,
	// 		&label.CreatedAt,
	// 		&label.UpdatedAt,
	// 	)
	// return label, err
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, name, number, user_id, created_at, updated_at from accounts where user_id=$1 limit 1`
	var model model_account.Account
	err := repo.DB.QueryRowContext(ctx, stmt, id).
		Scan(&model.ID, &model.Name, &model.Number, &model.UserId, &model.CreatedAt, &model.UpdatedAt)

	return &model, err
}

// List implements AccountRepository.
func (repo *accountRepository) List(id uuid.UUID) ([]model_account.Account, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, name, number, user_id, created_at, updated_at from accounts where user_id=$1`
	var resp = make([]model_account.Account, 0)
	rows, err := repo.DB.QueryContext(ctx, stmt, id)
	for rows.Next() {
		var model model_account.Account
		rows.Scan(&model.ID, &model.Name, &model.Number, &model.UserId, &model.CreatedAt, &model.UpdatedAt)
		resp = append(resp, model)
	}
	return resp, err
}
