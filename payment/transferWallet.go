package payment

// //request
type InitiateTransfer struct {
	Source    string `json:"source"`
	Reason    string `json:"reason"`
	Amount    int64  `json:"amount"`
	Reference string `json:"reference"`
	Recipient string `json:"recipient"`
}

// //webhook response
type TransferWebHook struct {
	Event string              `json:"event"`
	Data  DataWebhookTransfer `json:"data"`
}

type DataWebhookTransfer struct {
	Amount        int64       `json:"amount"`
	Currency      string      `json:"currency"`
	Domain        string      `json:"domain"`
	ID            int64       `json:"id"`
	Integration   Integration `json:"integration"`
	Reason        string      `json:"reason"`
	Reference     string      `json:"reference"`
	Source        string      `json:"source"`
	SourceDetails interface{} `json:"source_details"`
	Status        string      `json:"status"`
	TitanCode     interface{} `json:"titan_code"`
	TransferCode  string      `json:"transfer_code"`
	TransferredAt interface{} `json:"transferred_at"`
	Recipient     Recipient   `json:"recipient"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
}

type Integration struct {
	ID           int64  `json:"id"`
	IsLive       bool   `json:"is_live"`
	BusinessName string `json:"business_name"`
}

type Recipient struct {
	Active          bool        `json:"active"`
	Currency        string      `json:"currency"`
	Description     string      `json:"description"`
	Domain          string      `json:"domain"`
	Email           interface{} `json:"email"`
	ID              int64       `json:"id"`
	Integration     int64       `json:"integration"`
	Metadata        interface{} `json:"metadata"`
	Name            string      `json:"name"`
	RecipientCode   string      `json:"recipient_code"`
	Type            string      `json:"type"`
	IsDeleted       bool        `json:"is_deleted"`
	DetailsTransfer Details     `json:"details"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
}

type DetailsTransfer struct {
	AccountNumber string      `json:"account_number"`
	AccountName   interface{} `json:"account_name"`
	BankCode      string      `json:"bank_code"`
	BankName      string      `json:"bank_name"`
}

// ///initiatePayment Response

type TransferResponse struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Data    DataTransfer `json:"data"`
}

type DataTransfer struct {
	Domain        string      `json:"domain"`
	Amount        int64       `json:"amount"`
	Currency      string      `json:"currency"`
	Reference     string      `json:"reference"`
	Source        string      `json:"source"`
	SourceDetails interface{} `json:"source_details"`
	Reason        string      `json:"reason"`
	Status        string      `json:"status"`
	Failures      interface{} `json:"failures"`
	TransferCode  string      `json:"transfer_code"`
	TitanCode     interface{} `json:"titan_code"`
	TransferredAt interface{} `json:"transferred_at"`
	ID            int64       `json:"id"`
	Integration   int64       `json:"integration"`
	Recipient     int64       `json:"recipient"`
	CreatedAt     string      `json:"createdAt"`
	UpdatedAt     string      `json:"updatedAt"`
}
