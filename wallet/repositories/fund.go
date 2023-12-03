package repositories_wallet

import (
	"context"
	"personal-budget/util"
	wallet_model "personal-budget/wallet/model"

	"github.com/google/uuid"
)

// FundAccount implements WalletRepo.
func (repo *walletRepo) FundAccount(amount int, userId uuid.UUID) (wallet_model.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), util.DbTimeout)
	defer cancel()
	stmt := `update wallets set balance=balance+$1 where user_id=$2 returning id, balance, user_id, created_at, updated_at`
	var model wallet_model.Wallet
	err := repo.DB.QueryRowContext(ctx, stmt, amount, userId).
		Scan(&model.ID, &model.Balance, &model.UserID, &model.CreatedAt, &model.UpdatedAt)
	return model, err
}
