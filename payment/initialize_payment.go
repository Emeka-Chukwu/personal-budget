package payment

import (
	"fmt"
	"personal-budget/util"
)

func InitializePayment(config util.Config, model PayloadInit) (payload Payload, err error) {
	model.Amount = model.Amount + "00"
	model.Channels = []string{"card", "bank", "ussd", "qr", "mobile_money", "bank_transfer", "eft"}
	payload, err = HttpCaller[Payload]("post", fmt.Sprintf("%s/transaction/initialize", config.PaystackBaseURL), config.PaystackKey, model)
	return
}

func GetRecipientCode(config util.Config, model CreateRecipient) (payload RecipientResponse, err error) {
	payload, err = HttpCaller[RecipientResponse]("post", fmt.Sprintf("%s/transferrecipient", config.PaystackBaseURL), config.PaystackKey, model)
	return
}

func Create(config util.Config, model InitiateTransfer) (payload TransferResponse, err error) {
	payload, err = HttpCaller[TransferResponse]("post", fmt.Sprintf("%s/transfer", config.PaystackBaseURL), config.PaystackKey, model)
	return
}
