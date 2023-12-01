package repositories_users

import (
	"context"
	model_user "personal-budget/users/models"
	"personal-budget/util"
)

// GetUserByIDOrEmail implements UserAuthentication.
func (auth *authentication) GetUserByIDOrEmail(value any, key string) (model_user.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, first_name, last_name, email, phone, bvn, profile_url, created_at, updated_at from users where $1=$2`
	var user model_user.UserResponse
	err := auth.DB.QueryRowContext(ctx, stmt, key, value).
		Scan(
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
