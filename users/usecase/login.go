package usecase_user

import (
	"errors"
	"fmt"
	model_user "personal-budget/users/models"
	"personal-budget/util"
)

var (
	ErrInvalidEmailOrPassword = errors.New("invalid email or password")
	ErrValidateEmail          = errors.New("verify your email")
	ErrAccountSuspension      = errors.New("account is suspended")
)

// Login implements UsecaseUser.
func (us *usecaseuser) Login(data model_user.UserLogin) (model_user.UserRegisterResponse, error) {

	user, err := us.repo.GetUserByIDOrEmail(data.Email)
	if err != nil {
		return model_user.UserRegisterResponse{}, ErrInvalidEmailOrPassword
	}
	if err := util.CheckPassword(data.Password, user.Password); err != nil {
		return model_user.UserRegisterResponse{}, ErrInvalidEmailOrPassword
	}
	if !user.IsVerified {
		return model_user.UserRegisterResponse{}, ErrValidateEmail
	}
	if user.IsSuspended {
		return model_user.UserRegisterResponse{}, ErrAccountSuspension
	}
	access, payload, err := us.tokenMaker.CreateToken(user.ID, us.config.AccessTokenDuration, false)
	if err != nil {
		return model_user.UserRegisterResponse{}, err
	}
	refreshToken, payloadRefresh, err := us.tokenMaker.CreateToken(user.ID, us.config.RefreshTokenDuration, true)
	if err != nil {
		return model_user.UserRegisterResponse{}, err
	}
	//// set the values of session
	sess := model_user.Session{
		UserID:    user.ID,
		Refresh:   refreshToken,
		IsBlock:   false,
		UserAgent: "add the agent",
		ClientIP:  "add the client",
	}
	fmt.Println(sess)
	// create your wallet here
	// go func() {

	// }()
	resp := model_user.UserRegisterResponse{
		AccessToken:      access,
		RefreshToken:     refreshToken,
		AccessExpiredAt:  payload.ExpiredAt,
		AccessIssuedAt:   payload.IssuedAt,
		RefreshExpiredAt: payloadRefresh.ExpiredAt,
		RefreshIssuedAt:  payloadRefresh.IssuedAt,
		UserResponse:     user,
	}
	return resp, nil
}
