package repositories_users

import (
	"context"
	model_user "personal-budget/users/models"
	"personal-budget/util"

	"github.com/google/uuid"
)

// GetUserByIDOrEmail implements UserAuthentication.
func (auth *authentication) GetUserByEmail(email string) (model_user.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, first_name, last_name, email, phone, bvn, profile_url, created_at, updated_at, password, is_verified, is_suspended from users where email=$1`
	var user model_user.UserResponse
	err := auth.DB.QueryRowContext(ctx, stmt, email).
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
			&user.Password,
			&user.IsVerified,
			&user.IsSuspended,
		)
	return user, err
}

// GetUserByIDOrEmail implements UserAuthentication.
func (auth *authentication) GetUserById(id uuid.UUID) (model_user.UserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, first_name, last_name, email, phone, bvn, profile_url, created_at, updated_at, password, is_verified, is_suspended from users where id=$1`
	var user model_user.UserResponse
	err := auth.DB.QueryRowContext(ctx, stmt, id).
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
			&user.Password,
			&user.IsVerified,
			&user.IsSuspended,
		)
	return user, err
}
