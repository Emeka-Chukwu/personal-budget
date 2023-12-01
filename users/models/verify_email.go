package model_user

type VerifyEmail struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"secret_code" validate:"required,eq=6"`
}
