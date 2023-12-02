package repositories_users

import (
	"context"
	"personal-budget/util"

	"github.com/google/uuid"
)

// CreateSession implements UserAuthentication.
func (auth *authentication) CreateWallet(userId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `insert into wallets (user_id, balance)
		values ($1, $2 ) returning id, balance,user_id, created_at, updated_at`
	err := auth.DB.QueryRowContext(ctx, stmt,
		userId,
		0,
	).Err()
	return err
}
