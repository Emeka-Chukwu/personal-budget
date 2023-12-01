package repositories_users

import (
	"context"
	model_user "personal-budget/users/models"
	"personal-budget/util"

	"github.com/google/uuid"
)

// Update implements UserAuthentication.
func (auth *authentication) Update(id uuid.UUID, data model_user.User) (model_user.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update users set first_name=$1, last_name=$2, profile_url=$3, bvn=$4 where id=$5 returning id, first_name, last_name, email, phone, bvn, profile_url, created_at, updated_at`
	var user model_user.UserResponse
	err := auth.DB.QueryRowContext(ctx, stmt, data.FirstName, data.LastName, data.ProfileUrl, data.Bvn, id).
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
