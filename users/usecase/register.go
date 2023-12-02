package usecase_user

import (
	model_user "personal-budget/users/models"
	"personal-budget/util"

	"github.com/google/uuid"
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

	sess := model_user.Session{
		UserID:    user.ID,
		Refresh:   refreshToken,
		IsBlock:   false,
		UserAgent: data.UserAgent,
		ClientIP:  data.ClientIP,
	}
	var sessID uuid.UUID
	if sessRep, err := us.repo.CreateSession(sess); err == nil {
		sessID = sessRep.ID
	}

	go func() {
		us.repo.CreateWallet(user.ID)
	}()
	resp := model_user.UserRegisterResponse{
		AccessToken:      access,
		RefreshToken:     refreshToken,
		AccessExpiredAt:  payload.ExpiredAt,
		AccessIssuedAt:   payload.IssuedAt,
		RefreshExpiredAt: payloadRefresh.ExpiredAt,
		RefreshIssuedAt:  payloadRefresh.IssuedAt,
		UserResponse:     user,
		SessionId:        sessID,
	}
	return resp, nil
}
