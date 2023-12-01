package repositories_users

import (
	"context"
	model_user "personal-budget/users/models"
	"personal-budget/util"
)

// Register implements UserAuthentication.
func (auth *authentication) Register(data model_user.User) (model_user.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into users (first_name, last_name, email, phone, password, bvn, profile_url)
		values ($1, $2, $3, $4, $5, $6, $7) returning id, first_name, last_name, email, phone, bvn, profile_url, created_at, updated_at`
	var user model_user.UserResponse
	err := auth.DB.QueryRowContext(ctx, stmt,
		data.FirstName,
		data.LastName,
		data.Email,
		data.Phone,
		data.Password,
		data.Bvn,
		data.ProfileUrl,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Bvn,
		&user.ProfileUrl,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	return user, err

}
