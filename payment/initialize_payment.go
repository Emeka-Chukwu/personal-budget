package payment

import (
	"fmt"
	"personal-budget/util"
)

func (pay *paymentInterface) InitializePayment(model PayloadInit) (payload Payload, err error) {
	model.Amount = model.Amount + "00"
	model.Channels = []string{"card", "bank", "ussd", "qr", "mobile_money", "bank_transfer", "eft"}
	payload, err = HttpCaller[Payload]("post", fmt.Sprintf("%s/transaction/initialize", pay.config.PaystackBaseURL), pay.config.PaystackKey, model)
	return
}

func (pay *paymentInterface) GetRecipientCode(model CreateRecipient) (payload RecipientResponse, err error) {
	payload, err = HttpCaller[RecipientResponse]("post", fmt.Sprintf("%s/transferrecipient", pay.config.PaystackBaseURL), pay.config.PaystackKey, model)
	return
}

func (pay *paymentInterface) Create(model InitiateTransfer) (payload TransferResponse, err error) {
	payload, err = HttpCaller[TransferResponse]("post", fmt.Sprintf("%s/transfer", pay.config.PaystackBaseURL), pay.config.PaystackKey, model)
	return
}

type paymentInterface struct {
	config util.Config
}
type PaymentInterface interface {
	Create(model InitiateTransfer) (payload TransferResponse, err error)
	GetRecipientCode(model CreateRecipient) (payload RecipientResponse, err error)
	InitializePayment(model PayloadInit) (payload Payload, err error)
}

func NewPastackPayment(config util.Config) PaymentInterface {
	return &paymentInterface{config: config}
}
