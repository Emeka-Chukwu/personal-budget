package payment

type RecipientResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Data    DataResp `json:"data"`
}

type DataResp struct {
	Active        bool    `json:"active"`
	CreatedAt     string  `json:"createdAt"`
	Currency      string  `json:"currency"`
	Domain        string  `json:"domain"`
	ID            int64   `json:"id"`
	Integration   int64   `json:"integration"`
	Name          string  `json:"name"`
	RecipientCode string  `json:"recipient_code"`
	Type          string  `json:"type"`
	UpdatedAt     string  `json:"updatedAt"`
	IsDeleted     bool    `json:"is_deleted"`
	Details       Details `json:"details"`
}

type Details struct {
	AuthorizationCode interface{} `json:"authorization_code"`
	AccountNumber     string      `json:"account_number"`
	AccountName       interface{} `json:"account_name"`
	BankCode          string      `json:"bank_code"`
	BankName          string      `json:"bank_name"`
}

type CreateRecipient struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	Currency      string `json:"currency"`
}
