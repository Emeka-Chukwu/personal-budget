package usecase_user

import (
	"personal-budget/token"
	model_user "personal-budget/users/models"
	repositories_users "personal-budget/users/repositories"
	"personal-budget/util"
)

type UsecaseUser interface {
	Login(model_user.UserLogin) (model_user.UserRegisterResponse, error)
	Register(model_user.User) (model_user.UserRegisterResponse, error)
	VerifyEmail(data model_user.VerifyEmail) (model_user.UserResponse, error)
}

type usecaseuser struct {
	repo       repositories_users.UserAuthentication
	config     util.Config
	tokenMaker token.Maker
}

func NewUsecaseUser(config util.Config, token token.Maker, repo repositories_users.UserAuthentication) UsecaseUser {
	return &usecaseuser{repo: repo, config: config, tokenMaker: token}
}
