package payment

import "github.com/google/uuid"

type PayloadInit struct {
	Email    string   `json:"email"`
	Amount   string   `json:"amount"`
	Channels []string `json:"channels"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	Email  string    `json:"ticket_id"`
	UserID uuid.UUID `json:"user_id"`
}

type Payload struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}
