package usecase_user

import (
	model_user "personal-budget/users/models"
	"personal-budget/util"
)

// Register implements UsecaseUser.
func (us *usecaseuser) Register(data model_user.User) (model_user.UserRegisterResponse, error) {

	password, err := util.HashPassword(data.Password)
	if err != nil {
		return model_user.UserRegisterResponse{}, err
	}
	data.Password = password
	user, err := us.repo.Register(data)
	if err != nil {
		return model_user.UserRegisterResponse{}, err
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
	// Retrieve User-Agent from request headers

	sess := model_user.Session{
		UserID:    user.ID,
		Refresh:   refreshToken,
		IsBlock:   false,
		UserAgent: data.UserAgent,
		ClientIP:  data.ClientIP,
	}
	us.repo.CreateSession(sess)
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
