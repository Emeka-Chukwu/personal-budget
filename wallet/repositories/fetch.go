package repositories_wallet

import (
	"context"
	"personal-budget/util"
	wallet_model "personal-budget/wallet/model"

	"github.com/google/uuid"
)

// Fetch implements WalletRepo.
func (repo *walletRepo) Fetch(userId uuid.UUID) (wallet_model.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `select id, balance, user_id, created_at, updated_at from wallets where user_id=$1`
	var model wallet_model.Wallet
	err := repo.DB.QueryRowContext(ctx, stmt, userId).
		Scan(&model.ID, &model.Balance, &model.UserID, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}
