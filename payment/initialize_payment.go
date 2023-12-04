package payment

import "personal-budget/util"

func InitializePayment(config util.Config, model PayloadInit) (payload Payload, err error) {
	model.Amount = model.Amount + "00"
	model.Channels = []string{"card", "bank", "ussd", "qr", "mobile_money", "bank_transfer", "eft"}
	data, err := HttpCaller("post", "https://api.paystack.co/transaction/initialize", config.PaystackKey, model, &Payload{})
	if err != nil {
		return
	}
	payloadTemp := data.(*Payload)
	payload = *payloadTemp
	return
}
