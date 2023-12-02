package model_user

import (
	"personal-budget/shared"
	"time"

	"github.com/google/uuid"
)

type User struct {
	shared.Model
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required"`
	Password   string `json:"password" validate:"required,min=7"`
	Bvn        string `json:"bvn"`
	ProfileUrl string `json:"profile_url"`
	UserAgent  string
	ClientIP   string
}

type UserResponse struct {
	shared.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Bvn         string `json:"bvn"`
	ProfileUrl  string `json:"profile_url"`
	IsVerified  bool   `json:"is_verified"`
	IsSuspended bool   `json:"is_suspended"`
	Password    string `json:"password"`
}

type UserRegisterResponse struct {
	UserResponse
	AccessToken      string
	RefreshToken     string
	AccessExpiredAt  time.Time
	AccessIssuedAt   time.Time
	RefreshExpiredAt time.Time
	RefreshIssuedAt  time.Time
	SessionId        uuid.UUID
}

type UserLogin struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=7"`
	UserAgent string
	ClientIP  string
}
