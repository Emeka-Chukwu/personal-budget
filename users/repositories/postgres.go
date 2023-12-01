package repositories_users

import (
	"database/sql"
	model_user "personal-budget/users/models"

	"github.com/google/uuid"
)

type authentication struct {
	DB *sql.DB
}

func NewUserAuths(db *sql.DB) UserAuthentication {
	return &authentication{DB: db}

}

type UserAuthentication interface {
	// googleSignIn(data any)
	// faceboolSignIn(data any)
	// appleSignIn(data any)
	// twoFactorAuth(data any)

	Register(data model_user.User) (model_user.UserResponse, error)
	Update(id uuid.UUID, data model_user.User) (model_user.UserResponse, error)
	GetUserByIDOrEmail(id any, key string) (model_user.UserResponse, error)
}
