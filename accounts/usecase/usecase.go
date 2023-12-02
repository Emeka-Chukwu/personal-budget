package usecase_account

import (
	model_account "personal-budget/accounts/model"
	repositories_account "personal-budget/accounts/repositories"
	"personal-budget/token"
	"personal-budget/util"

	"github.com/google/uuid"
)

type AccountUsecase interface {
	Create(model model_account.Account) (*model_account.Account, error)
	Update(model model_account.Account, id uuid.UUID) (*model_account.Account, error)
	Delete(id uuid.UUID) error
	List(id uuid.UUID) ([]model_account.Account, error)
	Get(id uuid.UUID) (*model_account.Account, error)
	GetByID(id uuid.UUID) (*model_account.Account, error)
}

type accountUsecase struct {
	token  token.Maker
	repo   repositories_account.AccountRepository
	config util.Config
}

func NewAccountUsecase(token token.Maker, repo repositories_account.AccountRepository, config util.Config) AccountUsecase {
	return &accountUsecase{repo: repo, token: token, config: config}
}
