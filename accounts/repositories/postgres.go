package repositories_account

import (
	"database/sql"
	model_account "personal-budget/accounts/model"

	"github.com/google/uuid"
)

type AccountRepository interface {
	Create(model model_account.Account) (*model_account.Account, error)
	Update(model model_account.Account, id uuid.UUID) (*model_account.Account, error)
	Delete(userId uuid.UUID) error
	List(userId uuid.UUID) ([]model_account.Account, error)
	Get(userId uuid.UUID) (*model_account.Account, error)
	GetByID(id uuid.UUID) (*model_account.Account, error)
}

type accountRepository struct {
	DB *sql.DB
}

func NewAccountRepository(DB *sql.DB) AccountRepository {
	return &accountRepository{DB: DB}
}
