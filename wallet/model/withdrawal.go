package wallet_model

type WithdrawalRequest struct {
	ReceiveEmail string `json:"email"`
	Amount       int    `json:"amount"`
}
