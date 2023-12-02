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
	GetUserById(id uuid.UUID) (model_user.UserResponse, error)
	Register(data model_user.User) (model_user.UserResponse, error)
	CreateSession(data model_user.Session) (model_user.Session, error)
	Update(id uuid.UUID, data model_user.User) (model_user.UserResponse, error)
	GetUserByEmail(email string) (model_user.UserResponse, error)
	CreateWallet(userId uuid.UUID) error
}
