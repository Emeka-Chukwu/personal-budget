package payment

type ChargeWebhookResponse struct {
	Event string     `json:"event"`
	Data  DataCharge `json:"data"`
}

type DataCharge struct {
	ID                 int64       `json:"id"`
	Domain             string      `json:"domain"`
	Status             string      `json:"status"`
	Reference          string      `json:"reference"`
	Amount             int64       `json:"amount"`
	Message            interface{} `json:"message"`
	GatewayResponse    string      `json:"gateway_response"`
	DataPaidAt         string      `json:"paid_at"`
	CreatedAt          string      `json:"created_at"`
	Channel            string      `json:"channel"`
	Currency           string      `json:"currency"`
	IPAddress          string      `json:"ip_address"`
	Metadata           Metadata    `json:"metadata"`
	FeesBreakdown      interface{} `json:"fees_breakdown"`
	Log                interface{} `json:"log"`
	Fees               int64       `json:"fees"`
	FeesSplit          interface{} `json:"fees_split"`
	OrderID            interface{} `json:"order_id"`
	PaidAt             string      `json:"paidAt"`
	RequestedAmount    int64       `json:"requested_amount"`
	PosTransactionData interface{} `json:"pos_transaction_data"`
}
