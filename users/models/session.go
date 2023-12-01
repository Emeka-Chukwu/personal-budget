package model_user

import (
	"personal-budget/shared"

	"github.com/google/uuid"
)

type Session struct {
	shared.Model
	UserID    uuid.UUID `json:"user_id"`
	Refresh   string    `json:"refresh_token"`
	IsBlock   bool      `json:"is_block"`
	UserAgent string    `json:"user_agent"`
	ClientIP  string    `json:"client_ip"`
}
