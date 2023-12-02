package repositories_account

import (
	"database/sql"
	model_account "personal-budget/accounts/model"
)

type AccountRepository interface {
	Create(model model_account.Account) (*model_account.Account, error)
	Update(model model_account.Account) (*model_account.Account, error)
	Delete(model model_account.Account) error
	List(model model_account.Account) ([]model_account.Account, error)
	Get(model model_account.Account) (*model_account.Account, error)
}

type accountRepository struct {
	DB *sql.DB
}

// Create implements AccountRepository.
func (*accountRepository) Create(model model_account.Account) (*model_account.Account, error) {
	panic("unimplemented")
}

// Delete implements AccountRepository.
func (*accountRepository) Delete(model model_account.Account) error {
	panic("unimplemented")
}

// Get implements AccountRepository.
func (*accountRepository) Get(model model_account.Account) (*model_account.Account, error) {
	panic("unimplemented")
}

// List implements AccountRepository.
func (*accountRepository) List(model model_account.Account) ([]model_account.Account, error) {
	panic("unimplemented")
}

// Update implements AccountRepository.
func (*accountRepository) Update(model model_account.Account) (*model_account.Account, error) {
	panic("unimplemented")
}

func NewAccountRepository(DB *sql.DB) AccountRepository {
	return &accountRepository{DB: DB}
}
